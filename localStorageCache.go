package cache

import (
	utils "github.com/BabichMikhail/golang-dev-utils"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type LocalStorageCache struct {
	ICache

	baseDir    string
	timePrefix string
}

func NewLocalStorageCache(dir string) ICache {
	provider := new(LocalStorageCache)
	provider.baseDir = dir
	provider.timePrefix = "unix_time.expired_at:"
	return provider
}

func (p *LocalStorageCache) getFileName(key string) string {
	filename := utils.Sha512(key) + ".txt"
	return path.Join(p.baseDir, filename[0:2], filename[2:4], filename[4:6], filename[6:8], filename)
}

func (p *LocalStorageCache) isFileExists(filename string) bool {
	stat, err := os.Stat(filename)
	return !os.IsNotExist(err) && !stat.IsDir()
}

func (p *LocalStorageCache) Put(key string, content string, duration time.Duration) {
	filename := p.getFileName(key)
	utils.CheckNoError(os.MkdirAll(path.Dir(filename), 0755))
	utils.CheckNoError(ioutil.WriteFile(filename, []byte(strings.Join([]string{p.timePrefix + strconv.FormatInt(time.Now().Unix(), 10), content}, "\n")), 0755))
}

func (p *LocalStorageCache) Get(key string) (string, bool) {
	filename := p.getFileName(key)

	hasContent := false
	content := ""
	if p.isFileExists(filename) {
		rawContent := string(utils.CheckNoError(ioutil.ReadFile(filename)).([]byte))
		parts := strings.SplitN(rawContent, "\n", 1)
		if len(parts) == 2 {
			timeString := parts[0]
			ok := strings.HasPrefix(timeString, p.timePrefix) &&
				time.Now().Sub(time.Unix(utils.CheckNoError(strconv.ParseInt(strings.TrimPrefix(timeString, p.timePrefix), 10, 64)).(int64), 0)).Seconds() > 0
			if ok {
				hasContent = true
				content = parts[1]
			}
		}

		if !hasContent {
			p.Remove(key)
		}
	}

	return content, hasContent
}

func (p *LocalStorageCache) Remove(key string) {
	if filename := p.getFileName(key); p.isFileExists(filename) {
		utils.CheckNoError(os.Remove(filename))
	}
}
