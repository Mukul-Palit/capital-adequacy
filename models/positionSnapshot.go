package models

import (
	database "capital-adequacy/driver"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

//PositionSnapshot : Structure for the var table
type PositionSnapshot struct {
	ID        string  `gorm:"type:bigint(8);PRIMARY_KEY"`
	Timestamp string  `gorm:"type:datetime;NOT NULL"`
	Alpha     string  `gorm:"type:varchar(255);NOT NULL"`
	Days      string  `gorm:"type:varchar(255);NOT NULL"`
	Label     string  `gorm:"type:varchar(255);NOT NULL"`
	Exposure  float64 `gorm:"type:float;NOT NULL"`
	Var       float64 `gorm:"type:float;NOT NULL"`
}

//GetPositionSnapshot : returns 1 row that matches the condition
func GetPositionSnapshot() PositionSnapshot {
	var positionSnapshot PositionSnapshot
	db := database.DbConn()
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	var id, timestamp, alpha, days, label string
	var exposure, Var float64
	//fmt.Println(db.Table("var").Where("alpha = 0.99 AND days = 1 AND label = 'Warehouse-2011Parametric'").Order("ts desc").Take(&positionSnapshot).QueryExpr())
	//fmt.Println(db.Table("var").Order("id desc").Limit(1).Where("alpha = 0.99 AND days = 1 AND label = 'Warehouse-2011Parametric'").Select("id,ts,alpha,days,label,exposure,var").First(&positionSnapshot).QueryExpr())
	row := db.Table("var").Order("id desc").Limit(1).Where("alpha = 0.99 AND days = 1 AND label = 'Warehouse-2011Parametric'").Select("id,ts,alpha,days,label,exposure,var").First(&positionSnapshot).Row()
	row.Scan(&id, &timestamp, &alpha, &days, &label, &exposure, &Var)
	// fmt.Println(id, timestamp, alpha, days, label, exposure, Var)
	positionSnapshot = PositionSnapshot{ID: id, Timestamp: timestamp, Alpha: alpha, Days: days, Label: label, Exposure: exposure, Var: Var}
	//fmt.Println(positionSnapshot)
	return positionSnapshot
}
