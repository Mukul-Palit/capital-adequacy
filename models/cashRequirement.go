package models

import (
	database "capital-adequacy/driver"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

//CashRequirement : Structure for the cash_requirement table
type CashRequirement struct {
	ID              string  `gorm:"type:bigint(8);PRIMARY_KEY"`
	CashRequirement float64 `gorm:"type:decimal(12,2);NOT NULL"`
	MaxValueAtRisk  float64 `gorm:"type:decimal(12,2);NOT NULL"`
	PrimeBrokerCash float64 `gorm:"type:decimal(12,2);NOT NULL"`
	TotalCash       float64 `gorm:"type:decimal(12,2);NOT NULL"`
}

//GetCashRequirement : returns the latest entry in the table
func GetCashRequirement() CashRequirement {
	var cashRequirement CashRequirement
	db := database.DbConn()
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	//fmt.Println(db.Table("cash_requirement").Last(&cashRequirement).QueryExpr())
	db.Table("cash_requirement").Last(&cashRequirement)
	//fmt.Println(cashRequirement)
	return cashRequirement
}
