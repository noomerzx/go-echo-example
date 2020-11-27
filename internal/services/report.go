package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-api/internal/models"
	"os"
	"strconv"
)

// ReportService interface
type ReportService interface {
	GenReport(MerchantID int, Date string) error
}

// ReportServ struct
type ReportServ struct {
}

// NewReportService func
func NewReportService() *ReportServ {
	return &ReportServ{}
}

// GenReport func
func (rs ReportServ) GenReport(MerchantID int, Date string) error {
	fileName := strconv.Itoa(MerchantID) + ".json"
	f, err := os.Create("./report/" + fileName)
	if err != nil {
		panic(err)

	}
	defer f.Close()
	w := bufio.NewWriter(f)
	data := models.Report{
		Date: "27-11-2020",
		Products: []models.ProductSell{
			{
				Name:       "iPhone 12 mini",
				SellVolume: 1071,
			},
			{
				Name:       "iPhone 12",
				SellVolume: 3713,
			},
			{
				Name:       "iPhone 12 Pro",
				SellVolume: 1341,
			},
		},
		Accumulate: 230012003.00,
	}
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	n4, err := w.WriteString(string(b))
	if err != nil {
		panic(err)
	}
	w.Flush()
	fmt.Printf("wrote %d bytes\n", n4)
	return nil
}
