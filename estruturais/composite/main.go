package main

import (
	"github.com/masilvasql/composite/component"
	"github.com/masilvasql/composite/composite"
)

func main() {
	file1 := &component.File{Name: "file1.txt"}
	file2 := &component.File{Name: "file2.txt"}
	file3 := &component.File{Name: "file3.txt"}
	file4 := &component.File{Name: "file4.txt"}

	folder1 := &composite.Folder{Name: "folder1"}
	folder1.Add(file1)
	folder1.Add(file2)

	folder2 := &composite.Folder{Name: "folder2"}
	folder2.Add(file3)

	folder3 := &composite.Folder{Name: "folder3"}
	folder3.Add(file4)

	folder2.Add(folder3)

	root := &composite.Folder{Name: "root"}
	root.Add(folder1)
	root.Add(folder2)

	root.Show("")
}
