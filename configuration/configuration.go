package configuration

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/nanvenomous/exfs"
	"github.com/nanvenomous/ws/exception"
	"golang.org/x/mod/modfile"
)

const (
	GO_WORKFILE = "go.work"
	GO_MODFILE  = "go.mod"
)

var (
	errWorkFile error
)

func versionPreserver(path, version string) (string, error) { return version, nil }

type Configuration struct {
	FS           *exfs.FileSystem
	WorkFile     *modfile.WorkFile
	PathWorkFile string
}

func NewConfiguration(fs *exfs.FileSystem) *Configuration {
	var (
		cfg   *Configuration
		err   error
		bytes []byte
		mfl   *modfile.File
	)
	cfg = &Configuration{
		FS: fs,
	}
	cfg.PathWorkFile, err = fs.FindFileInAboveCurDir(GO_WORKFILE)
	if err != nil {
		errWorkFile = err
		return nil
	}

	bytes, err = ioutil.ReadFile(cfg.PathWorkFile)
	if err != nil {
		errWorkFile = err
		return nil
	}

	cfg.WorkFile, err = modfile.ParseWork(GO_WORKFILE, bytes, versionPreserver)
	if err != nil {
		errWorkFile = err
		return nil
	}

	err = os.Chdir(filepath.Dir(cfg.PathWorkFile))
	if err != nil {
		errWorkFile = err
		return nil
	}

	for i, wf := range cfg.WorkFile.Use {
		if wf.ModulePath == "" {
			bytes, err = ioutil.ReadFile(filepath.Join(wf.Path, GO_MODFILE))
			if err != nil {
				errWorkFile = err
				return nil
			}
			mfl, err = modfile.Parse(GO_MODFILE, bytes, versionPreserver)
			if err != nil {
				errWorkFile = err
				return nil
			}
			cfg.WorkFile.Use[i].ModulePath = mfl.Module.Mod.String()
		}
	}

	return cfg
}

func (cg *Configuration) Modules() []*modfile.Use {
	exception.CheckErr(errWorkFile)
	return cg.WorkFile.Use
}
