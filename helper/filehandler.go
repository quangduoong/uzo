package helper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func IsZipFile(src string) bool {
	splitFileName := strings.Split(src, ".zip")

	return len(splitFileName) == 2
}

func IsFileExists(src string) bool {
	_, err := os.Stat(src)

	return !os.IsNotExist(err) && err == nil
}

func IsDirExists(dir string) bool {
	_, err := os.Stat(dir)

	return !os.IsNotExist(err) && err == nil

}

func Unzip(src, dest string) error {
	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(file *zip.File) error {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, file.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), file.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := reader.Close(); err != nil {
			panic(err)
		}
	}()

	for _, f := range reader.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func OpenInExplorer(destPath string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	var dir string
	if filepath.IsAbs(destPath) {
		dir = destPath
	} else {
		dir = wd + "\\" + destPath
	}

	cmd := exec.Command("explorer", dir)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
