package models

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
)

type File struct {
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Content string `json:"content"`
}

type Directory struct {
	Path  string       `json:"path"`
	Dirs  *[]Directory `json:"dirs"`
	Files *[]File      `json:"files"`
}

type Node struct {
	Name     string
	Children []*Node
}

func NewDirectory(path string) Directory {
	newDir := Directory{}
	newDir.Path = path
	newDir.Dirs = &[]Directory{}
	newDir.Files = &[]File{}
	return newDir
}

func DirsExplorer(rootDir *Directory, recursiveScan bool, index *Index) {
	var wg sync.WaitGroup
	wg.Add(1)
	go dirsRunner(rootDir, recursiveScan, &wg, index)
	wg.Wait()
}

func dirsRunner(rootDir *Directory, recursiveScan bool, wg *sync.WaitGroup, index *Index) {
	defer wg.Done()
	files, err := ioutil.ReadDir(rootDir.Path)

	if err != nil {
		fmt.Println("Error dirsRunner:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			newDir := NewDirectory(filepath.Join(rootDir.Path, file.Name()))
			*rootDir.Dirs = append(*rootDir.Dirs, newDir)
			if !recursiveScan {
				wg.Add(1)
				go dirsRunner(&newDir, recursiveScan, wg, index)
			}
		} else if !file.IsDir() {

			fileURL := filepath.Join(rootDir.Path, file.Name())

			messageID := ReadMail(fileURL, "message-id")
			from := ReadMail(fileURL, "from")
			to := ReadMail(fileURL, "to")
			content := ReadMail(fileURL, "content")

			*rootDir.Files = append(*rootDir.Files, File{file.Name(), file.Size(), ""})

			*index.Records = append(*index.Records, Mail{messageID, from, to, content})

		}
	}

}
