package models

type Login struct {
	Email string `json:"email,omitempty" validate:"nonzero,matches=/^((?!\.)[\w-_.]*[^.])(@\w+)(\.\w+(\.\w+)?[^.\W])$/gim;`
	Senha string `json:"senha,omitempty" validate:"nonzero,min=8"`
}
