package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddFont("PFBeauSansPro-Light", "", "./fonts/PFBeauSansPro-Light.json")
	pdf.AddFont("PFBeauSansPro-Book", "", "./fonts/PFBeauSansPro-Book.json")
	pdf.AddFont("PFBeauSansPro-Bold", "", "./fonts/PFBeauSansPro-Bold.json")
	pdf.AddPage()
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")

	pdf.Image("images/body.png", float64(0), float64(0), float64(297), float64(210), false, "", 0, "")
	pdf.SetFont("PFBeauSansPro-Light", "", 28)
	pdf.Cellf(0, 0, tr("СЕРТИФИКАТ"))
	pdf.SetFont("PFBeauSansPro-book", "", 29)
	pdf.SetXY(float64(19), float64(120))
	pdf.Cellf(156, 25, tr("Начертательная геометрия и инженерная графика"))
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(float64(19), float64(147))
	pdf.Cellf(100, 30, tr("3 зачетные единицы"))
	err := pdf.OutputFileAndClose("results/pdf.pdf")
	if err != nil {
		log.Println(err)
	}
}
