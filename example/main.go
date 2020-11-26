package main

import (
	"fmt"
	"log"

	u "github.com/c-seeger/Golang-HTML-TO-PDF-Converter"
)

func main() {

	r := u.NewRequestPdf("")

	//html template path
	templatePath := "sample.html"

	//path for download pdf
	outputPath := "example.pdf"

	//html template data
	templateData := struct {
		Title       string
		Description string
		Company     string
		Contact     string
		Country     string
	}{
		Title:       "HTML to PDF generator",
		Description: "This is the simple HTML to PDF file.",
		Company:     "Jhon Lewis",
		Contact:     "Maria Anders",
		Country:     "Germany",
	}

	if err := r.ParseTemplate(templatePath, templateData); err != nil {
		log.Fatal(err)
	}
	if err := r.GeneratePDF(outputPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pdf generated successfully")
}
