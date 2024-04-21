package service

import (
	"bytes"
	"fmt"
	"generate-report-pdf/model"
	"log"
	"strings"
	"text/template"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

var (
	templatePath = "./template/report-product.html"
	filePath     = "./output/report-product.pdf"
)

func InitData() map[string]interface{} {
	productInfo := model.ProductInfo{
		Id:          "PSerial01",
		Name:        "Product 1",
		Description: "Product 1 description",
	}
	inventoryList := []model.Inventory{
		{Location: "D1", Quantity: 1000, Available: 900},
		{Location: "D2", Quantity: 200, Available: 50},
		{Location: "D3", Quantity: 700, Available: 700},
		{Location: "D4", Quantity: 1000, Available: 1},
		{Location: "D5", Quantity: 1000, Available: 1},
		{Location: "D6", Quantity: 1000, Available: 1},
		{Location: "D7", Quantity: 1000, Available: 200},
		{Location: "D8", Quantity: 1000, Available: 70},
		{Location: "D9", Quantity: 1000, Available: 30},
		{Location: "D10", Quantity: 1000, Available: 0},
		{Location: "D11", Quantity: 1000, Available: 100},
	}
	saleList := []model.Sale{
		{Date: "15/04/2024", QuantitySold: 500, TotalRevenue: 500000000},
		{Date: "16/04/2024", QuantitySold: 300, TotalRevenue: 300000000},
		{Date: "17/04/2024", QuantitySold: 110000, TotalRevenue: 11000000000},
		{Date: "18/04/2024", QuantitySold: 110000, TotalRevenue: 11000000000},
		{Date: "19/04/2024", QuantitySold: 110000, TotalRevenue: 11000000000},
		{Date: "20/04/2024", QuantitySold: 110000, TotalRevenue: 11000000000},
		{Date: "21/04/2024", QuantitySold: 110000, TotalRevenue: 11000000000},
	}

	data := map[string]interface{}{
		"productName": productInfo.Name,
		"productId":   productInfo.Id,
		"description": productInfo.Description,
	}

	var inventoryStr string
	for _, in := range inventoryList {
		inventoryStr += fmt.Sprintf(`<tr>
	   <td>%s</td>
	   <td>%d</td>
	   <td>%d</td>
	 </tr>`, in.Location, in.Quantity, in.Available)
	}
	data["inventory"] = inventoryStr

	var saleStr string
	for _, in := range saleList {
		saleStr += fmt.Sprintf(`<tr>
	   <td>%s</td>
	   <td>%d</td>
	   <td>%f</td>
	 </tr>`, in.Date, in.QuantitySold, in.TotalRevenue)
	}
	data["sale"] = saleStr

	return data
}

func InitPDF() *wkhtmltopdf.PDFGenerator {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	return pdfg
}

func GenerateReport() error {
	// Init data
	data := InitData()

	// Init PDF page setting
	pdf := InitPDF()

	// Read the HTML template file
	htmlTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	// Execute the template and write the result to a string
	htmlContent := new(strings.Builder)
	err = htmlTemplate.Execute(htmlContent, data)
	if err != nil {
		return err
	}

	// Convert string to []byte
	htmlBytes := []byte(htmlContent.String())
	reader := bytes.NewReader(htmlBytes)

	// Create a new PDF document
	pdf.AddPage(wkhtmltopdf.NewPageReader(reader))

	// Create PDF document in internal buffer
	err = pdf.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write pdf file
	pdf.WriteFile(filePath)

	return nil
}
