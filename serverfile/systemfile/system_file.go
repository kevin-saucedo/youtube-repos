package systemfile

import (
	"context"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
)

type FileSystem interface {
	Close() error
	Upload(filekey string, data []byte) error
	Delete(fileky string) error
	Serve(rw http.ResponseWriter, filekey string, name string) error
}

type system struct {
	ctx    context.Context
	bucket *blob.Bucket
}

func NewFileSystem(dirpath string, ctx ...context.Context) (FileSystem, error) {
	c := context.Background()
	if len(ctx) > 0 {
		c = ctx[0]
	}
	if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
		return nil, err
	}
	bucket, err := fileblob.OpenBucket(dirpath, nil)
	if err != nil {
		return nil, err
	}
	return &system{ctx: c, bucket: bucket}, nil
}

func (s *system) Close() error {
	return s.bucket.Close()
}

func (s *system) Upload(filekey string, data []byte) error {
	opts := &blob.WriterOptions{
		ContentType: mimetype.Detect(data).String(),
	}
	w, err := s.bucket.NewWriter(s.ctx, filekey, opts)
	if err != nil {
		return err
	}
	defer w.Close()
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}
func (s *system) Delete(filekey string) error {
	return s.bucket.Delete(s.ctx, filekey)
}

var inlineServeContentTypes = []string{
	"image/png", "image/jpg", "image/jpeg", "image/gif", "image/webp", "image/x-icon", "image/bmp",
	"video/webm", "video/mp4", "video/3gpp", "video/quicktime", "video/x-ms-wmv",
	"audio/basic", "audio/aiff", "audio/mpeg", "audio/midi", "audio/mp3", "audio/wave",
	"audio/wav", "audio/x-wav", "audio/x-mpeg", "audio/x-m4a", "audio/aac",
	"application/pdf", "application/x-pdf",
}

func (s *system) Serve(rw http.ResponseWriter, filekey string, name string) error {
	r, err := s.bucket.NewReader(s.ctx, filekey, nil)
	if err != nil {
		return err
	}
	defer r.Close()
	disposition := "attachment" // archivo se va a descargar
	contentType := r.ContentType()
	if ExistInSlice(contentType, inlineServeContentTypes) {
		disposition = "inline" // se a mostar en el navegador
	}
	rw.Header().Del("X-Frame-Options")
	rw.Header().Set("Content-Disposition", disposition+"; filename="+name)
	rw.Header().Set("Content-Type", contentType)
	rw.Header().Set("Content-Length", strconv.FormatInt(r.Size(), 10))
	rw.Header().Set("Content-Security-Policy", "default-src 'none'; media-src 'self'; style-src 'unsafe-inline'; sandbox")
	location, err := time.LoadLocation("GMT")
	if err != nil {
		rw.Header().Set("Last-Modified", r.ModTime().In(location).Format("Mon, 02 Jan 06 15:04:05 MST"))
	}
	if _, err := io.Copy(rw, r); err != nil {
		return err
	}
	return nil
}

func ExistInSlice[T comparable](item T, list []T) bool {
	if len(list) == 0 {
		return false
	}
	for index := range list {
		if list[index] == item {
			return true
		}
	}
	return false
}
