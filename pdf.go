package pdfgenerator

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

//RequestPdf contains the html data
type RequestPdf struct {
	Body       string
	localFiles bool
}

// LocalFileAccess enables or disables local file access
func (r *RequestPdf) LocalFileAccess(b bool) {
	r.localFiles = b
}

//NewRequestPdf creates a new RequestPdf from body
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		Body: body,
	}
}

//ParseTemplate to template data
func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return nil
}

//GeneratePDF generates the pdf from the request
func (r *RequestPdf) GeneratePDF(pdfPath string) error {
	f, err := ioutil.TempFile("", "html2pdf*.html")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())

	if _, err := f.WriteString(r.Body); err != nil {
		f.Close()
		return err
	}
	f.Close()

	// super strange bug have to reopen the file again
	f, err = os.Open(f.Name())
	if err != nil {
		return err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	page := wkhtmltopdf.NewPageReader(f)
	if r.localFiles {
		page.EnableLocalFileAccess.Set(true)
	}
	pdfg.AddPage(page)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
