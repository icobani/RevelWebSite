package model

import (
	"encoding/json"
	"strings"
)

type QuickSettingUp struct {
	Id     int64 `json:"-"`
	UserId int64 `json:"-"`

	Step int `json:"step"`

	Company           string `json:"company_name" CaptionML:"enu=Company Name;trk=Firma Adı"`
	Name              string `json:"name" CaptionML:"enu=Name;trk=Surname"`
	Email             string `json:"email" CaptionML:"enu=E-Mail;trk=E-Mail"`
	Country						string	` CaptionML:"enu=Country;trk=Ülke"`
	PhoneNumber       string `json:"phone_number" CaptionML:"enu=Phone Number;trk=Telefon Numarası"`
	ExpenseCategories string `json:"expense_categories" CaptionML:"enu=Expense Categories;trk=Masraf Kategorileriniz"`
	Branches          string `json:"branches" CaptionML:"enu=Branches;trk=Şubeler"`
	Projects          string `json:"projects" CaptionML:"enu=Projects;trk=Projeler"`
	Departments       string `json:"departments" CaptionML:"enu=Departments;trk=Departmanlar"`
	Invites           string `json:"stuff_emails" CaptionML:"enu=Invites;trk=Davetler"`
	Password					string	` CaptionML:"enu=Password;trk=Şifre"`
}

type JQuickSettingUp QuickSettingUp

func (q QuickSettingUp) MarshalJSON() ([]byte, error) {
	if q.Id == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JQuickSettingUp(q))
}

func (q *QuickSettingUp) GetExpenseCategories() []string {
	if q.ExpenseCategories == "" {
		return []string{"Default"}
	}
	return strings.Split(q.ExpenseCategories, ",")
}

func (q *QuickSettingUp) GetBranches() []string {
	if q.Branches == "" {
		return []string{"Default"}
	}
	return strings.Split(q.Branches, ",")
}

func (q *QuickSettingUp) GetProjects() []string {
	if q.Projects == "" {
		return []string{"Default"}
	}
	return strings.Split(q.Projects, ",")
}

func (q *QuickSettingUp) GetDepartments() []string {
	if q.Departments == "" {
		return []string{"Default"}
	}
	return strings.Split(q.Departments, ",")
}

func (q *QuickSettingUp) GetInvites() []string {
	if q.Invites == "" {
		return []string{}
	}
	return strings.Split(q.Invites, ",")
}
