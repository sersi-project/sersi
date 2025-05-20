package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateDirectory(name string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	projectPath := filepath.Join(cwd, name)
	err = os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Project Created at %s\n", projectPath)
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
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
