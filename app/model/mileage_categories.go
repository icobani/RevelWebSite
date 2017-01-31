package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type MileageCategory struct {
	Id           int64      `json:"id" CaptionML:"enu=ID;trk=ID"`
	CompanyId    int64      `json:"-" CaptionML:"enu=Company;trk=Şirket"`
	BranchId     int64      `json:"-" CaptionML:"enu=Branch;trk=Şube"`
	DepartmentId int64      `json:"-" CaptionML:"enu=Department;trk=Departman"`
	Name         string     `json:"name" sql:"type:varchar(250);" CaptionML:"enu=Mileage Category Name;trk=Mil Kategori Adı"`
	RatePerKm    float32    `json:"rate_per_km" CaptionML:"enu=Rate Per Km;trk=Oran/Km"`
	CurrencyCode string     `json:"currency" sql:"type:varchar(50);" CaptionML:"enu=Currency Code;trk=Döviz Kodu"`
	StartDate    *time.Time `json:"start_date,omitempty" CaptionML:"enu=Start Date;trk=Başlangıç Tarihi"`
	IsActive     bool       `json:"is_active" CaptionML:"enu=Active;trk=Aktif"`
}

func (this MileageCategory) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("MileageCategory Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("MileageCategory Table Created")
	app.MakeCaptionML(this)
}