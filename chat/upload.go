package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadrHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userid")
	file, header, err := r.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	filename := filepath.Join("avatars", userId+filepath.Ext(header.Filename))
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "Successful")
}
