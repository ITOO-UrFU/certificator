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
	//pdf.Image("images/template-03.png", float64(0), float64(0), float64(297), float64(210), false, "", 0, "")
	//pdf.Line(17, 0, 17, 300)
	pdf.SetFont("PFBeauSansPro-Light", "", 28)
	pdf.SetTextColor(111, 112, 114)
	pdf.SetXY(17, 82.7)
	pdf.Cellf(0, 0, tr("СЕРТИФИКАТ"))
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(79.5, 84.3)
	pdf.Cellf(0, 0, tr("подтверждает, что"))
	pdf.SetXY(17, 112.377)
	pdf.Cellf(0, 0, tr("успешно освоил(-а) курс"))
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(17, 158.058)
	pdf.MultiCell(180, 6, tr("Описание освоенного курса и достигнутых результатов обучения приведено\nв приложении к настоящему сертификату"), "", "", false)
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(17, 182.793)
	pdf.Cellf(0, 0, tr("Электронная версия сертификата:"))
	pdf.SetXY(145, 182.793)
	pdf.Cellf(0, 0, tr("Сертификат №"))
	pdf.SetXY(145, 188.714)
	pdf.Cellf(0, 0, tr("выдан"))
	pdf.SetXY(217, 182.793)
	pdf.Cellf(0, 0, tr("Зам. проректора УрФУ"))

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("PFBeauSansPro-Bold", "", 38)
	pdf.SetXY(17, 98.62)
	pdf.Cellf(0, 0, tr("Имя Фамилия Отчество"))
	pdf.SetFont("PFBeauSansPro-book", "", 29)
	pdf.SetXY(17, 120.161)
	pdf.MultiCell(180, 12, tr(strings.ToUpper("Начертательная геометрия \nи инженерная графика")), "", "", false)
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(17, 149.952)
	pdf.Cellf(0, 0, tr("3 зачетные единицы"))
	pdf.SetFont("PFBeauSansPro-Light", "", 14)
	pdf.SetXY(17, 188.714)
	pdf.Cellf(0, 0, "https://openedu.urfu.ru/certificates/")
	pdf.SetXY(217, 188.714)
	pdf.Cellf(0, 0, tr("Третьяков Василий Сергеевич"))

	//x, y := pdf.GetPageSize()
	//pdf.AddPageFormat("P", gofpdf.SizeType{y, x})
	err := pdf.OutputFileAndClose("results/pdf.pdf")
	if err != nil {
		log.Println(err)
	}
}
