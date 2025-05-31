package utils

import (
	"embed"
	"io"
	"io/fs"
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

func CopyDirectory(src embed.FS,folder string, dst string) error {
	dstPath := filepath.Join(dst)

	err := fs.WalkDir(src, folder, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(folder, path)
		if err != nil {
			return err
		}

		dstFilePath := filepath.Join(dstPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstFilePath, os.ModePerm)
		}

		content, err := src.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(dstFilePath, content, os.ModePerm)
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
