package extensions

import (
	"path"

	"github.com/charmbracelet/log"
	"github.com/go-git/go-git/v5"
	"github.com/yuin/gopher-lua"
	"selene.frankmayer.dev/util"
)

// Clones a git repository to a specified destination. If the repository already exists, it will pull the latest changes instead.
// returns true if successful, false otherwise
func gitCloneOrPull(l *lua.LState) int {
	url := l.CheckString(1)
	dest := path.Join(util.ConfigPath, l.CheckString(2))

	var err error
	var repo *git.Repository

	// check if repo exists
	repo, err = git.PlainOpen(dest)
	if err != nil {
		// repo does not exist, clone it
		log.Debug("Cloning", "repo", url, "dest", dest)
		repo, err = git.PlainClone(dest, false, &git.CloneOptions{
			URL: url,
		})
		if err != nil {
			log.Error("Error cloning repo", "error", err)
			l.Push(lua.LFalse)
			return 1
		}
		l.Push(lua.LTrue)
		return 1
	}

	// repo exists, pull latest changes
	wt, err := repo.Worktree()
	if err != nil {
		log.Error("Error pulling repo", "error", err)
		l.Push(lua.LFalse)
		return 1
	}
	err = wt.Pull(&git.PullOptions{})
	if err != nil {
		if err == git.NoErrAlreadyUpToDate {
			log.Debug("Repo already up to date", "repo", url, "dest", dest)
			l.Push(lua.LTrue)
			return 1
		}
		log.Error("Error pulling repo", "error", err)
		l.Push(lua.LFalse)
		return 1
	}

	l.Push(lua.LTrue)
	return 1
}
