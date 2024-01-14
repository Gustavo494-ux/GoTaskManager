package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nome       string `json:"nome,omitempty" validate:"min=5"`
	CPF        string `json:"cpf,omitempty" validate:"nonzero,regexp=^[0-9]{11}$"`
	Email      string `json:"email,omitempty" validate:"nonzero,matches=/^((?!\.)[\w-_.]*[^.])(@\w+)(\.\w+(\.\w+)?[^.\W])$/gim;`
	Email_Hash string `json:"Email_Hash,omitempty" serializar:"false"`
	Senha      string `json:"senha,omitempty" validate:"nonzero,min=8" serializar:"false"`
}
