package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

//Rate : Structure for the symbol exposure table
type Rate struct {
	ID           string `gorm:"type:bigint;PRIMARY_KEY"`
	Date         string `gorm:"type:date;NOT NULL"`
	Data         []byte `gorm:"type:varchar(255);NOT NULL"`
	BaseCurrency string `gorm:"type:varchar(255);NOT NULL"`
	CreatedAt    string `gorm:"type:datetime;NOT NULL"`
}
