package main

import (
	"fmt"
	"hamocha/cmd"
)

func main() {
	dir, _ := cmd.ParseDirectory()
	files, _ := cmd.GetDirectoryEntries(dir)
	categorizedFiles, _ := cmd.CategorizeFiles(files)
	fmt.Println(categorizedFiles)
	cmd.CreateOrganizingDirectory(categorizedFiles)
	cmd.MoveFilesToCategoryDirectory(files, dir)
}
