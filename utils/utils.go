package utils

import (
	"io"
	"os"
	"path/filepath"
)

func CreateDirectory(name string) error {
	projectPath := GetProjectPath(name)
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func CopyDirectory(src, dst string) error {
	srcPath := filepath.Join(src)
	dstPath := filepath.Join(dst)

	err := os.MkdirAll(dstPath, os.ModePerm)
	if err != nil {
		return err
	}

	err = filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(srcPath, path)
		dstFilePath := filepath.Join(dstPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstFilePath, os.ModePerm)
		}

		return CopyFile(path, dstFilePath, info)
	})

	if err != nil {
		return err
	}

	return nil
}

func CopyFile(src, dst string, info os.FileInfo) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := in.Close(); cerr != nil {
			err = cerr
		}
	}()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}

	defer func() {
		if cerr := out.Close(); cerr != nil {
			err = cerr
		}
	}()

	_, err = io.Copy(out, in)
	return err
}

func FileExists(name string) bool {
	path := GetProjectPath(name)
	_, err := os.Stat(path)
	return err == nil
}

func GetProjectPath(name string) string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(cwd, name)
}
