package main

import (
	"fmt"
	"io"
	"net/http"
	"serverfile/systemfile"

	"github.com/rs/xid"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/upload", uploadfilehandle)
	http.HandleFunc("/download", downloadfilehandle)
	fmt.Println("0.0.0.0:3003")
	http.ListenAndServe("0.0.0.0:3003", nil)
}

func uploadfilehandle(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		responseError(w, err)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		responseError(w, err)
		return
	}
	storer, err := systemfile.NewFileSystem("files")
	if err != nil {
		responseError(w, err)
		return
	}
	defer storer.Close()

	filekey := xid.New().String()
	if err := storer.Upload(filekey, data); err != nil {
		responseError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(filekey))
}
func downloadfilehandle(w http.ResponseWriter, r *http.Request) {
	filekey := r.URL.Query().Get("filekey")
	storer, err := systemfile.NewFileSystem("files")
	if err != nil {
		responseError(w, err)
		return
	}
	defer storer.Close()
	if err := storer.Serve(w, filekey, "myfile"); err != nil {
		responseError(w, err)
	}
}
func responseError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
