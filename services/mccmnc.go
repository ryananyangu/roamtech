package services

import (
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/ryananyangu/roamtech/database"
	"github.com/ryananyangu/roamtech/models"
	"github.com/ryananyangu/roamtech/utils"
)

const COLUMN_SIZE = 7

func ProcessMccMnc() ([]models.MccMnc, error) {

	res, err := utils.Request("", map[string][]string{}, "https://www.mcc-mnc.com/", "GET")
	if err != nil {
		utils.ErrorLogger.Println(err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		utils.ErrorLogger.Println(err)
		return nil, err
	}

	return HtmlToObjList(doc), nil

}

func HtmlToObjList(doc *goquery.Document) []models.MccMnc {

	list := []models.MccMnc{}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {

		row := ""

		s.Find("td").Each(func(v int, x *goquery.Selection) {

			row += x.Contents().Text() + "^"

		})

		items := strings.Split(row, "^")
		if len(items) == COLUMN_SIZE {

			obj := models.MccMnc{
				MCC:         items[0],
				MNC:         items[1],
				ISO:         items[2],
				Country:     items[3],
				CountryCode: items[4],
				Network:     items[5],
			}

			SaveToDatabase(&obj)

			list = append(list, obj)
		}

	})

	return list

}

func SaveToDatabase(mcc_mnc *models.MccMnc) *models.MccMnc {

	networkcode := mcc_mnc.MCC + "-" + mcc_mnc.MNC
	db := *database.GetDB()
	db[networkcode] = *mcc_mnc

	return mcc_mnc

}

func GetByCountry(country string) []models.MccMnc {

	list := []models.MccMnc{}

	db := *database.GetDB()

	for _, obj := range db {
		if strings.EqualFold(country, obj.Country) {
			list = append(list, obj)
		}

	}

	return list

}

func GetByMcc(mcc string) []models.MccMnc {

	db := *database.GetDB()

	list := []models.MccMnc{}

	for _, obj := range db {
		// Check if the current country is same as the one searched

		if mcc == obj.MCC {

			list = append(list, obj)
		}
	}

	return list

}
