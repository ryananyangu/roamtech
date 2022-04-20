package database

import "github.com/ryananyangu/roamtech/models"

var Database map[string]models.MccMnc

func init() {

	Database = map[string]models.MccMnc{}

}

func GetDB() *map[string]models.MccMnc {
	return &Database

}
