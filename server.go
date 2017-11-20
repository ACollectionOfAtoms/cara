package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
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
		// TODO: do this and facial processing in memory!
		saveFile(fileName, r)
	}
	fmt.Println("files saved.")
	fmt.Println("Processing images...")
	processFaces()
	// fmt.Println(output)
	// if e != nil {
	// 	fmt.Println("error!")
	// 	fmt.Println(e)
	// } else {
	// 	fmt.Println("image processing complete")
	// 	fmt.Println(output)
	// }
}

func processFaces() {
	cmd := exec.Command("./compare.py", "test/*")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	go copyOutput(stdout)
	go copyOutput(stderr)
	cmd.Wait()
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
