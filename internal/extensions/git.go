package extensions

import (
	"github.com/Frank-Mayer/yab/internal/util"

	"errors"
	"path"

	"github.com/charmbracelet/log"
	"github.com/go-git/go-git/v5"
	"github.com/yuin/gopher-lua"
)

// Clones a git repository to a specified destination. If the repository already exists, it will pull the latest changes instead.
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
            l.Error(lua.LString("Error cloning repo. " + err.Error()), 0)
			return 0
		}
		return 0
	}

	// repo exists, pull latest changes
	wt, err := repo.Worktree()
	if err != nil {
        l.Error(lua.LString("Error getting worktree. " + err.Error()), 0)
        return 0
	}
	err = wt.Pull(&git.PullOptions{})
	if err != nil {
		if errors.Is(err, git.NoErrAlreadyUpToDate) {
			log.Debug("Repo already up to date", "repo", url, "dest", dest)
			return 0
		}
		l.Error(lua.LString("Error pulling repo. " + err.Error()), 0)
		return 0
	}

	return 0
}
