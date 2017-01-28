package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type MileageCategory struct {
	Id           int64      `json:"id"`
	CompanyId    int64      `json:"-"`
	BranchId     int64      `json:"-"`
	DepartmentId int64      `json:"-"`
	Name         string     `json:"name" sql:"type:varchar(250);"`
	RatePerKm    float32    `json:"rate_per_km"`
	Currency     string     `json:"currency" sql:"type:varchar(50);`
	StartDate    *time.Time `json:"start_date,omitempty"`
	IsActive     bool       `json:"is_active"`
}

func (this MileageCategory) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("MileageCategory Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("MileageCategory Table Created")
}