package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("index.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        data := map[string]interface{}{
            "Title": "Hello, World!",
        }
        tmpl.Execute(w, data)
    })

    http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
        // Ensure the method is POST
        if r.Method != http.MethodPost {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        // Parse the form to handle file upload
        err := r.ParseMultipartForm(10 << 20) // Limit file size to 10 MB
        if err != nil {
            http.Error(w, "Unable to parse form", http.StatusBadRequest)
            return
        }

        // Get the uploaded file
        file, _, err := r.FormFile("source_input")
        if err != nil {
            http.Error(w, "Unable to get file", http.StatusBadRequest)
            return
        }

        defer file.Close()

        contents, errs := io.ReadAll(file)

        if errs != nil {
            http.Error(w, "Failed to read file", http.StatusInternalServerError)
            return 
        }
        
         fileContent := string(contents)

        
        
        tmpl, err := template.ParseFiles("index.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        data := map[string]interface{}{
            "Title": "Hello, World!",
            "output": func() string {
                return fileContent
            }(),
        }
        tmpl.Execute(w, data)

        // Respond to the client (optional: can be an HTMX response)
    })

    http.ListenAndServe(":8080", nil)
    fmt.Println("Server started on http://localhost:8080")
}