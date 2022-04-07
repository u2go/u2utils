package u2utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
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
	dir, err := filepath.Abs(dir)
	if err != nil {
		return "", nil
	}
	for i := 0; i < 100; i++ {
		filename1, err := filepath.Abs(path.Join(dir, filename))
		if err != nil {
			return "", err
		}
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

func FileListAll(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			files = append(files, strings.ReplaceAll(path, "\\", "/"))
			return nil
		})
	if err != nil {
		return nil, err
	}
	return files, nil
}
