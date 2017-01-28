package model

import (
	"time"
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)

type Export struct {
	Id      int64  `json:"id"`
	Filter  string `json:"filter" sql:"type:text;"`
	Name    string `json:"name" sql:"type:varchar(255);"`
	Status  string `json:"status" sql:"type:varchar(20);"`
	Comment string `json:"comment" sql:"type:text;`
	Address string `json:"address" sql:"type:varchar(255);"`

	CreatedAt time.Time `json:"created_at"`
}

func (this Export) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("Export Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("Export Table Created")
}