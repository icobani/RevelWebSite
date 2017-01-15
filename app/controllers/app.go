package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {
	return c.Render()
}
func (c App) Pricing() revel.Result {
	return c.Render()
}
func (c App) Tour() revel.Result {
	return c.Render()
}
func (c App) Us() revel.Result {
	return c.Render()
}
func (c App) Why() revel.Result {
	return c.Render()
}