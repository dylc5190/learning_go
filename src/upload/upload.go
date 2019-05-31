package main

import (
    "io"
    "log"
    "path/filepath"
    "net/http"
    "os"
    //"strings"
)

const (
    DefaultUploadDir = "."   
)

func ReceiveFormFile(w http.ResponseWriter, r *http.Request) {
    // if err = req.ParseMultipartForm(2 << 10); err != nil {  
    //    status = http.StatusInternalServerError  
    //    return  
    // }  
    
    // r.Method should be "POST"
    file, header, err := r.FormFile("file")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    //nameParts := strings.Split(header.Filename, ".")
    //filename := nameParts[1]
    savedPath := filepath.Join(DefaultUploadDir, header.Filename)
    f, err := os.OpenFile(savedPath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    _, err = io.Copy(f, file)
    if err != nil {
        panic(err)
    }

    return
}

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/upload", ReceiveFormFile)
    log.Fatal(http.ListenAndServe(":8080", nil))
}