package pdf

import (
	"fmt"
	"log"
	"strconv"

	"movido-media/controllers/billing"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var currencyMap = map[string]string{
	"EUR": "€",
	"USD": "$",
}

func currencySymbol(symbol string, price float64) string {
	if val, ok := currencyMap[symbol]; ok {
		return fmt.Sprintf("%s %.2f", val, price)
	}

	return fmt.Sprintf("%s %.2f", symbol, price)
}

func (ip invoicePDF) Generate(details billing.ContractDetail) (string, error) {
	m := ip.NewMarto(details)
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	filename := fmt.Sprintf("reports/%s.pdf", details.ProductCode)

	if err = document.Save(filename); err != nil {
		return "", err
	}

	return filename, nil
}

func (ip invoicePDF) NewMarto(details billing.ContractDetail) core.Maroto {
	cfg := config.NewBuilder().
		WithMargins(10, 15, 10).
		WithPageSize(pagesize.A4).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	if err := m.RegisterHeader(ip.pageHeader()); err != nil {
		log.Fatal(err.Error())
	}

	if err := m.RegisterFooter(ip.pageFooter()); err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(text.NewRow(10, "Invoice ABC123456789", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Align: align.Center,
	}))

	m.AddRow(7,
		text.NewCol(3, "Transactions", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: ip.setColor(50, 205, 50)})

	m.AddRows(ip.getTable(details)...)
	return m
}

func (ip invoicePDF) setColor(r, b, g int) *props.Color {
	return &props.Color{
		Red:   r,
		Blue:  b,
		Green: g,
	}
}

func (ip invoicePDF) pageHeader() core.Row {
	return row.New(20).Add(
		image.NewFromFileCol(3, "assets/logo.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		col.New(2),
		col.New(6).Add(
			text.New("Movido Media Verlag GmbH",
				props.Text{
					Size:  16,
					Align: align.Center,
					Color: ip.setColor(0, 0, 100),
				}),
			text.New("Steinstrasse 2 | 40212 Düsseldorf", props.Text{
				Top:   12,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Center,
				Color: ip.setColor(0, 100, 0),
			}),
			text.New("www.movido-media.de", props.Text{
				Top:   15,
				Style: fontstyle.BoldItalic,
				Size:  6,
				Align: align.Center,
				Color: ip.setColor(0, 100, 0),
			}),
		),
	)
}

func (ip invoicePDF) getTable(details billing.ContractDetail) []core.Row {
	rows := []core.Row{
		row.New(5).Add(
			text.NewCol(4, "Product Code", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(2, "Product Name", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Price", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Billed For (Months)", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		).WithStyle(&props.Cell{BackgroundColor: ip.setColor(201, 201, 201)}),
	}

	var contentsRow []core.Row

	r := row.New(4).Add(
		text.NewCol(4, details.ProductCode, props.Text{Size: 8, Align: align.Center}),
		text.NewCol(2, details.ProductName, props.Text{Size: 8, Align: align.Center}),
		text.NewCol(3, currencySymbol(details.Currency, details.Price), props.Text{Size: 8, Align: align.Center}),
		text.NewCol(2, strconv.Itoa(details.BillingFrequency), props.Text{Size: 8, Align: align.Center}),
	)

	contentsRow = append(contentsRow, r)

	rows = append(rows, contentsRow...)

	rows = append(rows, row.New(20).Add(
		col.New(7),
		text.NewCol(2, "Total:", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(3, currencySymbol(details.Currency, details.Price*float64(details.BillingFrequency)), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Center,
		}),
	))

	return rows
}

func (ip invoicePDF) pageFooter() core.Row {
	return row.New(20).Add(
		col.New(12).Add(
			text.New("Tel:  +49 (0) 211 205 470", props.Text{
				Top:   13,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: ip.setColor(0, 100, 0),
			}),
			text.New("www.movido-media.de", props.Text{
				Top:   16,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: ip.setColor(0, 100, 0),
			}),
		),
	)
}
