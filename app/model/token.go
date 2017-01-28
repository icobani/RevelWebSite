package model

import (
	"fmt"
	"github.com/icobani/RevelWebSite/app"
)


type TokenStorage struct {
	Id   string `json:"email"`
	Hash string `json:"hash"`
}

func (this TokenStorage) CreateTable() {
	app.DB.DropTable(this)
	fmt.Println("TokenStorage Table Dropped")
	app.DB.CreateTable(this)
	fmt.Println("TokenStorage Table Created")
}