package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// echo "# HTMX" >> README.md
// git init
// git add README.md
// git commit -m "first commit"
// git branch -M main
// git remote add origin git@github.com:vigneshwaran-ravi/HTMX.git
// git push -u origin main
func main()  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "index.html")
		tmpl,err := template.ParseFiles("index.html");
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := map[string]interface{}{
			"Title": "Hello, World!",
		}
		// data := struct{
		// 	message string
		// }{
		// 	message: "hello world",
		// }
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("hello world")
}