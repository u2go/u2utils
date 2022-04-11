package u2utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

var tmpDir = path.Join(os.TempDir(), "u2-tmp-cache")

func getFile(key string) string {
	return path.Join(tmpDir, SHA256(key))
}

func TmpCacheGet(key string, valP any) error {
	file := getFile(key)
	exists, err := FileExists(file)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
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
