package main

import (
	"embed"
	"io/fs"
)

var corpusFs embed.FS

func SetFs(fs embed.FS) {
	corpusFs = fs
}

func GetFs() embed.FS {
	return corpusFs
}

func GetFileSystem(p ...string) fs.FS {
	path := "corpus"
	if len(p) != 0 && p[0] != "" {
		path = p[0]
	}
	fsys, _ := fs.Sub(GetFs(), path)
	return fsys
}
