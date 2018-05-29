package local

import (
	"github.com/usmanhalalit/gost"
)

type LocalFilesystem struct {
	Service interface{}
	Config  Config
}

type Config struct {
	BasePath string
}

func New(c Config) gost.Directory {
	fs := LocalFilesystem{
		Service: nil,
		Config: c,
	}
	return &Directory{
		Object{
			Fs:   &fs,
			Path: c.BasePath,
		},
	}
}

func (fs *LocalFilesystem) GetClient() interface{} {
	return fs.Service
}

func (fs *LocalFilesystem) GetConfig() interface{} {
	return fs.Config
}