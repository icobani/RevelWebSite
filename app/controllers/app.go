package controllers

import (
	"github.com/revel/revel"
	"fmt"
	"net/http"
	"github.com/icobani/RevelWebSite/app"
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

func (c App) QuickRegister() revel.Result {
	fmt.Println("QuickRegister....")
	return c.Render()
}

func (c App) QuickRegister_POST() revel.Result {
	fmt.Println("QuickRegister_POST")
	return c.RenderTemplate("Login/Index.html")
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