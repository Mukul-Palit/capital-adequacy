package models

import (
	database "airflow-report/capital-adequacy/driver"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

type Pos interface {
	GetPositionSnapshot() (PositionSnapshot, error)
}

type Position struct {
	DB *gorm.DB
}

//PositionSnapshot : Structure for the var table
type PositionSnapshot struct {
	ID        string  `gorm:"type:bigint(8);PRIMARY_KEY" json:"id"`
	Timestamp string  `gorm:"type:datetime;NOT NULL" json:"ts"`
	Alpha     string  `gorm:"type:varchar(255);NOT NULL" json:"alpha"`
	Days      string  `gorm:"type:varchar(255);NOT NULL" json:"days"`
	Label     string  `gorm:"type:varchar(255);NOT NULL" json:"label"`
	Exposure  float64 `gorm:"type:float;NOT NULL" json:"exposure"`
	Var       float64 `gorm:"type:float;NOT NULL" json:"Var"`
}

//GetPositionSnapshot : returns 1 row that matches the condition
func (p *Position) GetPositionSnapshot() (PositionSnapshot, error) {
	var positionSnapshot PositionSnapshot
	var id, timestamp, alpha, days, label string
	var exposure, Var float64
	row, err := p.DB.Table("var").Order("id desc").Limit(1).Where("alpha = 0.99 AND days = 1 AND label = 'Warehouse-2011Parametric'").Select("id,ts,alpha,days,label,exposure,var").Rows()
	defer row.Close()
	if err != nil {
		database.WriteLogFile(err)
	}
	for row.Next() {
		row.Scan(&id, &timestamp, &alpha, &days, &label, &exposure, &Var)
	}
	positionSnapshot = PositionSnapshot{ID: id, Timestamp: timestamp, Alpha: alpha, Days: days, Label: label, Exposure: exposure, Var: Var}
	return positionSnapshot, err
}

func CreatePositionSnapshot(db *gorm.DB) Pos {
	return &Position{
		DB: db,
	}
}
