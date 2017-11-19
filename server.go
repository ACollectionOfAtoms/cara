package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func saveFile(fileName string, r *http.Request) {
	file, _, err := r.FormFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	filepath := fmt.Sprintf("./test/%v", fileName)

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("GET used, implement error here!")
	} else {
		fmt.Println("Receiving Files...")
	}
	r.ParseMultipartForm(32 << 20)
	files := r.MultipartForm.File

	fmt.Printf("Saving %v files...\n", len(files))
	for fileName := range files {
		saveFile(fileName, r)
	}
	fmt.Print("files saved.")
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
