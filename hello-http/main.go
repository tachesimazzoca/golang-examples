package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type MimeType struct {
	Suffix      string
	ContentType string
}

var mimeTypes = []MimeType{
	MimeType{".html", "text/html"},
	MimeType{".css", "text/css"},
	MimeType{".js", "text/javascript"},
	MimeType{".jpg", "image/jpeg"},
	MimeType{".gif", "image/gif"},
	MimeType{".png", "image/png"},
	MimeType{".xml", "application/xml"},
	MimeType{".json", "application/json"},
	MimeType{".txt", "text/plain"},
}

func main() {

	fgs := flag.NewFlagSet("server", flag.ExitOnError)
	docRoot := fgs.String("d", ".", "document root")
	port := fgs.Int("p", 4000, "port")
	if err := fgs.Parse(os.Args); err != nil {
		log.Fatal(err)
	}
	fgs.Parse(os.Args[1:])

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		doc := path.Join(*docRoot, req.RequestURI)
		if body, e := ioutil.ReadFile(doc); e == nil {
			rw.Header().Set("Content-Type", "application/octet-stream")
			for _, mt := range mimeTypes {
				if strings.HasSuffix(doc, mt.Suffix) {
					log.Println(mt)
					rw.Header().Set("Content-Type", mt.ContentType)
					break
				}
			}
			rw.WriteHeader(http.StatusOK)
			io.WriteString(rw, string(body))
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	})
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: withLogging(mux),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func withLogging(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			log.Println(req.RequestURI)
			handler.ServeHTTP(rw, req)
		})
}
