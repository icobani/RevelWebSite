package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"net/http"
)

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {
	fmt.Println("Incoming Index")
	return c.Render()
}

func (c App) Pricing() revel.Result {
	fmt.Println("Incoming Pricing")
	return c.Render()
}

func (c App) Tour() revel.Result {
	return c.Render()
}

func (c App) Us() revel.Result {
	return c.Render()
}

func (c App) Why() revel.Result {
	fmt.Println("Why....")
	return c.Render()
}

func (c App) Pet() revel.Result {
	fmt.Println("Pet....")
	return c.Render()
}

func (c App) Met() revel.Result {
	fmt.Println("Met....")
	return c.Render()
}

// Set for Turkish Lang
func (c App)SetLang_TR() revel.Result {
	new_cookie := &http.Cookie{Name: "REVEL_LANG", Value: "tr"}
	c.SetCookie(new_cookie)
	// Return the refered page
	return c.Redirect(c.Request.Referer())
}

//Set for English Lang
func (c App)SetLang_EN() revel.Result {
	new_cookie := &http.Cookie{Name: "REVEL_LANG", Value: "en"}
	c.SetCookie(new_cookie)
	// Return the refered page
	return c.Redirect(c.Request.Referer())
}