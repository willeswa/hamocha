package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ParseDirectory() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("directory argument is required")
	}

	return os.Args[1], nil
}

func GetDirectoryEntries(dir string) ([]File, error) {
	directories, err := os.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	var files []File

	for _, d := range directories {
		info, err := d.Info()

		if err != nil {
			return nil, err
		}

		if !info.IsDir() {
			file := ExtractFileInfo(info)
			files = append(files, file)
		}
	}

	return files, nil
}

func ExtractFileInfo(fileInfo os.FileInfo) File {
	fileExt := strings.ToLower(filepath.Ext(fileInfo.Name()))
	file := File{
		Name:  fileInfo.Name(),
		Size:  int(fileInfo.Size()),
		IsDir: fileInfo.IsDir(),
		Ext:   fileExt,
	}
	return file
}

func GetFileCategory(fileExt string) string {
	for category, exts := range fileTypes {
		for _, ext := range exts {
			if ext == fileExt {
				return category
			}
		}
	}

	return "Other"
}

func CategorizeFiles(files []File) (map[string][]File, error) {

	categorizedFiles := make(map[string][]File)

	for _, file := range files {
		category := GetFileCategory(file.Ext)
		categorizedFiles[category] = append(categorizedFiles[category], file)
	}
	return categorizedFiles, nil
}

func CreateOrganizingDirectory(categorizedFiles map[string][]File) (bool, error) {

	workingDir, err := os.Getwd()

	if err != nil {
		return false, err
	}

	for category := range categorizedFiles {
		folderPath := filepath.Join(workingDir, category)

		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			if err := os.Mkdir(folderPath, 0755); err != nil {
				return false, err
			}

		}
	}
	return true, nil
}

func MoveFilesToCategoryDirectory(files []File, sourceDir string) error {
	workingdir, err := os.Getwd()

	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir {
			continue
		}


		category := GetFileCategory(file.Ext)
		source := filepath.Join(workingdir, file.Name)
		fmt.Println(source)
		dest := filepath.Join(workingdir, category, file.Name)
		fmt.Println(dest)
		if err := os.Rename(source, dest); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
