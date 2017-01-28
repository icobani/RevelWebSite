package model

import (
	"encoding/json"

	"github.com/guregu/null"
)

// S(Simplified) Model Structs

type SUser struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type SProject struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type SGroup struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type SDepartment struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type JSDepartment SDepartment

func (d SDepartment) MarshalJSON() ([]byte, error) {
	if d.Id.Int64 == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JSDepartment(d))
}

type SBranch struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type JSBranch SBranch

func (b SBranch) MarshalJSON() ([]byte, error) {
	if b.Id.Int64 == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JSBranch(b))
}

type SExpenseCategory struct {
	Id   null.Int    `json:"id"`
	Name null.String `json:"name"`
}

type JSExpenseCategory SExpenseCategory

func (ec SExpenseCategory) MarshalJSON() ([]byte, error) {
	if ec.Id.Int64 == 0 {
		return []byte("null"), nil
	}
	return json.Marshal(JSExpenseCategory(ec))
}
