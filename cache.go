package u2utils

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"syscall"
)

var tmpDir = path.Join(os.TempDir(), "u2-tmp-cache")

func getFile(key string) string {
	return path.Join(tmpDir, SHA256(key))
}

func TmpCacheGet(key string, valP any) error {
	file := getFile(key)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		if err1, ok := err.(*fs.PathError); ok &&
			err1.Err == syscall.ERROR_FILE_NOT_FOUND {
			return nil
		}
		return err
	}
	err = json.Unmarshal(b, valP)
	if err != nil {
		return err
	}
	return nil
}

func TmpCacheSet(key string, value any) error {
	err := FileMkdirAll(tmpDir)
	if err != nil {
		return err
	}
	file := getFile(key)
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, v, 0777)
}
