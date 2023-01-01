package models

import (

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