package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func base64Image(fileName string, r *http.Request) string {
	file, _, err := r.FormFile(fileName)
	if err != nil {
		panic("Failed to read file!")
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Failed to read file!")
	}
	b64String := base64.RawStdEncoding.EncodeToString(fileBytes)
	defer file.Close()
	return b64String
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("POST was not used, implement error here!")
	} else {
		fmt.Println("Receiving Files...")
	}
	r.ParseMultipartForm(32 << 20)
	files := r.MultipartForm.File

	fmt.Printf("Saving %v files...\n", len(files))

	imageStrings := make([]string, 0, 2)
	fmt.Println("Converting Image to base64 string")
	for fileName := range files {
		b64String := base64Image(fileName, r)
		imageStrings = append(imageStrings, b64String)
	}
	fmt.Println("files converted to base64 strings.")

	fmt.Println("Processing images...")
	res, err := processFaces(r, imageStrings[0], imageStrings[1])
	if err != nil {
		fmt.Println(err)
		fmt.Println("ERROR")
	}
	fmt.Println("Succesfully Processed Image!")

	fmt.Println("Writing data to response...")
	fmt.Fprintf(w, "%s", res)
	fmt.Println("Response sent.")
}

func processFaces(r *http.Request, imageString1, imageString2 string) (string, error) {
	pythonProcess := exec.Command("./compare.py")

	stdin, err := pythonProcess.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}
	stdout, err := pythonProcess.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdin.Close()
	pythonProcess.Stderr = os.Stderr

	fmt.Println("Starting comparison...")
	if err = pythonProcess.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}

	io.WriteString(stdin, imageString1)
	io.WriteString(stdin, "\n")
	io.WriteString(stdin, imageString2)
	io.WriteString(stdin, "\n")
	results, err := ioutil.ReadAll(stdout)
	resultsString := string(results)
	if err != nil {
		panic(err)
	}
	resultSlice := strings.Split(resultsString, "!")
	fmt.Println(resultSlice)
	percent := resultSlice[len(resultSlice)-1]
	err = pythonProcess.Wait()
	if err != nil {
		panic(err)
	}
	return percent, err
}

func stringFromIOReader(r io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/upload", upload)
	log.Fatal(http.ListenAndServe(":7777", nil))
}
