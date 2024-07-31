package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
)

const MAX_UPLOAD_SIZE = 1024 * 1024

func main() {
	http.HandleFunc("GET /", homeView)
	http.HandleFunc("GET /files", fileView)
	http.HandleFunc("POST /files/upload", fileUpload)
	http.HandleFunc("POST /files/download", fileDownload)

	fmt.Println("server running on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func homeView(w http.ResponseWriter, r *http.Request) {
	templ := `
		<html>
			<head>
				<title>Home | App</title>
			</head>
			<body>
				<h1>Welcome to the App</h1>
				Go to <a href="/files">files page</a>
			</body>
		</html>
	`

	fmt.Fprint(w, templ)
}

func fileView(w http.ResponseWriter, r *http.Request) {
	templ := `
		<html>
			<head>
				<title>Home | App</title>
			</head>
			<body>
				<h2>Upload File</h2>
				<form method="POST" action="/files/upload" enctype="multipart/form-data">
					<input type="file" name="myFile" />
					<button>Upload</button>
				</form>
				<h2>Download File</h2>
				<form method="POST" action="/files/download">
					<input type="text" name="myFile">
					<button>Download</button>
				</form>
			</body>
		</html>	
	`

	fmt.Fprint(w, templ)
}

func fileUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Println("File name:", fileHeader.Filename)
	fmt.Println("Header:", fileHeader.Header)
	fmt.Printf("Size: %d bytes\n", fileHeader.Size)

	// Create the 'uploads' folder if it doesn't already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the 'uploads' folder
	dst, err := os.Create(fmt.Sprintf("./uploads/%s%s", uuid.NewString(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/files", http.StatusSeeOther)
}

func fileDownload(w http.ResponseWriter, r *http.Request) {
	filename := r.FormValue("myFile")

	file, err := os.Open(fmt.Sprintf("./uploads/%s", filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer file.Close()

	tempBuffer := make([]byte, 512)
	file.Read(tempBuffer)
	fileContentType := http.DetectContentType(tempBuffer)

	fileStat, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileSize := strconv.FormatInt(fileStat.Size(), 10)

	w.Header().Set("Content-Type", fileContentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Length", fileSize)

	file.Seek(0, 0)
	io.Copy(w, file)
}
