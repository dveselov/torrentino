package main

import (
	"fmt"
	"mime"
	"path/filepath"
)

type File struct {
	Length   int64  `json:"length"`
	Path     string `json:"path"`
	Mimetype string `json:"mimetype"`
	URL      string `json:"url"`
}

func GetMimeTypeByPath(path string) string {
	ext := filepath.Ext(path)
	mimetype := mime.TypeByExtension(ext)
	return mimetype
}

func SerializeTorrentFiles(client *Client) []File {
	originalFiles := client.Torrent.Files()
	files := make([]File, 0)
	for i := 0; i < len(originalFiles); i++ {
		file := originalFiles[i]
		length, path := file.Length(), file.Path()
		var protocol string
		if TorrentinoHTTPS {
			protocol = "https"
		} else {
			protocol = "http"
		}
		if length > 0 {
			mimetype := GetMimeTypeByPath(path)
			token, _ := NewToken(client.Config.URL, path)
			files = append(files, File{
				length,
				path,
				mimetype,
				fmt.Sprintf("%s://%s/%s/", protocol, TorrentinoHostname, token),
			})
		}
	}
	return files
}
