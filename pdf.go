package main

import (
	"log"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddFont("PFBeauSansPro-Light", "", "./fonts/PFBeauSansPro-Light.json")
	pdf.AddFont("PFBeauSansPro-Book", "", "./fonts/PFBeauSansPro-Book.json")
	pdf.AddFont("PFBeauSansPro-Bold", "", "./fonts/PFBeauSansPro-Bold.json")
	pdf.AddPage()
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")

	pdf.Image("images/background.png", float64(0), float64(0), float64(297), float64(210), false, "", 0, "")
	pdf.Image("images/signature.png", float64(0), float64(0), float64(297), float64(210), false, "", 0, "")
	pdf.Image("images/body.png", float64(0), float64(0), float64(297), float64(210), false, "", 0, "")
	pdf.SetFont("PFBeauSansPro-Light", "", 28)
	pdf.SetXY(17, 84.5)
	pdf.Cellf(0, 0, tr("СЕРТИФИКАТ"))
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(79.5, 86)
	pdf.Cellf(0, 0, tr("подтверждает, что"))
	pdf.SetXY(17, 117.5)
	pdf.Cellf(0, 0, tr("успешно освоил(-а) курс"))
	pdf.SetFont("PFBeauSansPro-book", "", 29)
	pdf.SetXY(17, 130)
	pdf.MultiCell(180, 10, tr(strings.ToUpper("Начертательная геометрия \nи инженерная графика")), "1", "", false)
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(17, 167)
	pdf.Cellf(0, 0, tr("3 зачетные единицы"))
	err := pdf.OutputFileAndClose("results/pdf.pdf")
	if err != nil {
		log.Println(err)
	}
}
