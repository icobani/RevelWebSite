package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)


type TokenStorage struct {
	Id   string `json:"email" CaptionML:"enu=ID;trk=ID"`
	Hash string `json:"hash" CaptionML:"enu=Hash;trk=Bilet"`
}

func (this TokenStorage) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("TokenStorage Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("TokenStorage Table Created")
	app.MakeCaptionML(this)
}