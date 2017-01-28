package model

import (
	"encoding/json"
	"strings"
)

type QuickSettingUp struct {
	Id     int64 `json:"-"`
	UserId int64 `json:"-"`

	Step int `json:"step"`

	Company           string `json:"company_name"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	ExpenseCategories string `json:"expense_categories"`
	Branches          string `json:"branches"`
	Projects          string `json:"projects"`
	Departments       string `json:"departments"`
	Invites           string `json:"stuff_emails"`
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
		return []string{}
	}
	return strings.Split(q.ExpenseCategories, ",")
}

func (q *QuickSettingUp) GetBranches() []string {
	if q.Branches == "" {
		return []string{}
	}
	return strings.Split(q.Branches, ",")
}

func (q *QuickSettingUp) GetProjects() []string {
	if q.Projects == "" {
		return []string{}
	}
	return strings.Split(q.Projects, ",")
}

func (q *QuickSettingUp) GetDepartments() []string {
	if q.Departments == "" {
		return []string{}
	}
	return strings.Split(q.Departments, ",")
}

func (q *QuickSettingUp) GetInvites() []string {
	if q.Invites == "" {
		return []string{}
	}
	return strings.Split(q.Invites, ",")
}
