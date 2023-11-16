package extensions

import (
	"archive/zip"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/yuin/gopher-lua"
)

var (
    // array of filenames to exclude from zip
    zipBlacklist = []string{
        ".DS_Store",
        "thumbs.db",
    }
)

// Create a zip file containing the given files. Returns true if successful, false otherwise.
func makeZip(l *lua.LState) int {
	files_t := l.CheckTable(1)
	output := l.CheckString(2)

	archive, err := os.Create(output)
	if err != nil {
		log.Error(err)
		l.Push(lua.LFalse)
		return 1
	}
	defer archive.Close()

	writer := zip.NewWriter(archive)
	defer writer.Close()

	files_t.ForEach(func(_ lua.LValue, value lua.LValue) {
		path := value.String()

        fi, err := os.Stat(path)
        if os.IsNotExist(err) {
            log.Error("File does not exist", "path", path)
            return // continue
        }

        if fi.IsDir() {
		if err := addFilesToZip(writer, path, path); err != nil {
			log.Error(err)
		}
        } else {
			dat, err := os.ReadFile(path)
            if err != nil {
                log.Error(err)
                return // continue
            }

            f, err := writer.Create(fi.Name())
            if err != nil {
                log.Error(err)
                return // continue
            }
            _, err = f.Write(dat)
            if err != nil {
                log.Error(err)
                return // continue
            }
        }
	})

	l.Push(lua.LTrue)
	return 1
}

func addFilesToZip(w *zip.Writer, basePath, baseInZip string) error {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}

    files_loop:
	for _, file := range files {
        // check if file is in blacklist
        for _, blacklisted := range zipBlacklist {
            if file.Name() == blacklisted {
                continue files_loop
            }
        }

		fullfilepath := filepath.Join(basePath, file.Name())
        fi, err := os.Stat(fullfilepath)
		if os.IsNotExist(err) {
			continue
		}

		if fi.Mode()&(fs.ModeIrregular|fs.ModeSymlink|fs.ModeDevice|fs.ModeNamedPipe|fs.ModeSocket) != 0 {
			// skip irregular files (e.g. symlinks)
            log.Warn("Skipping irregular file", "path", fullfilepath)
			continue
		} else if fi.IsDir() {
			if err := addFilesToZip(w, fullfilepath, filepath.Join(baseInZip, file.Name())); err != nil {
				return err
			}
		} else {
            in_zip_path := filepath.Join(baseInZip, file.Name())
			dat, err := os.ReadFile(fullfilepath)
			if err != nil {
				return err
			}

			f, err := w.Create(in_zip_path)
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
