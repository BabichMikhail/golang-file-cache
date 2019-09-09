package cache

import (
	utils "github.com/BabichMikhail/golang-dev-utils"
	"io/ioutil"
	"os"
	"path"
)

type LocalStorageCache struct {
	ICache

	baseDir string
}

func NewLocalStorageCache(dir string) ICache {
	provider := new(LocalStorageCache)
	provider.baseDir = dir
	return provider
}

func (p *LocalStorageCache) getFileName(key string) string {
	filename := utils.Md5(key) + ".txt"
	return path.Join(p.baseDir, "cache", filename[0:2], filename[2:4], filename[4:6], filename[6:8], filename)
}

func (p *LocalStorageCache) Put(key string, content string) {
	filename := p.getFileName(key)
	utils.CheckNoError(os.MkdirAll(path.Dir(filename), 0755))
	utils.CheckNoError(ioutil.WriteFile(filename, []byte(content), 0755))
}

func (p *LocalStorageCache) Get(key string) (string, bool) {
	filename := p.getFileName(key)
	stat, err := os.Stat(filename)

	fileExists := !os.IsNotExist(err) && !stat.IsDir()
	content := ""
	if fileExists {
		content = string(utils.CheckNoError(ioutil.ReadFile(filename)).([]byte))
	}

	return content, fileExists
}
