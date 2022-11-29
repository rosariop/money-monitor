package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("server starting")
	http.HandleFunc("/hello", fileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mediaType)

	mr := multipart.NewReader(r.Body, params["boundary"])
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		slurp, err := io.ReadAll(p)
		if err != nil {
			log.Fatal(err)
		}
		auszug := string(slurp)

		// Splitting CSV
		rows := strings.Split(auszug, "\n")
		rows = KillEmptyRows(rows)

		for i, row := range rows {
			strings.Replace(row, "\r", "", -1)
			fmt.Printf("-------%d-------\n", i)
			cells := strings.Split(row, ";")
			for _, cell := range cells {
				fmt.Printf("%s;", cell)
			}
		}

	}

	w.Write([]byte("hello world"))
}
