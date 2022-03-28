package u2utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

func FileExists(filename string) (bool, error) {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func FileExistsParent(dir, filename string) (string, error) {
	for i := 0; i < 100; i++ {
		filename1 := path.Join(dir, filename)
		exists, err := FileExists(filename1)
		// 出错了
		if err != nil {
			return "", err
		}
		// 存在了
		if exists {
			return filename1, nil
		}
		dir1 := filepath.Dir(dir)
		if dir1 == dir {
			return "", errors.New("no file found until root dir")
		}
		dir = dir1
	}
	return "", errors.New("no file found until max loop")
}
