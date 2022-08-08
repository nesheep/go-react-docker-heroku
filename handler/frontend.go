package handler

import (
	"embed"
	"errors"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"
)

type FrontEnd struct {
	assets embed.FS
	prefix string
}

func NewFrontend(assets embed.FS, prefix string) *FrontEnd {
	return &FrontEnd{
		assets: assets,
		prefix: prefix,
	}
}

func (f *FrontEnd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f.tryRead(r.URL.Path, w)
	if err == nil {
		return
	}

	err = f.tryRead("index.html", w)
	if err != nil {
		panic(err)
	}
}

func (f *FrontEnd) tryRead(requestPath string, w http.ResponseWriter) error {
	file, err := f.assets.Open(path.Join(f.prefix, requestPath))
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return errors.New("path is dir")
	}

	ext := filepath.Ext(requestPath)
	var contentType string
	if m := mime.TypeByExtension(ext); m != "" {
		contentType = m
	} else {
		contentType = "application/octet-stream"
	}

	w.Header().Set("Content-Type", contentType)
	io.Copy(w, file)

	return nil
}
