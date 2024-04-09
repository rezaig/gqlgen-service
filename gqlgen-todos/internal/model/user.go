package model

type User struct {
	ID string `json:"id"`
	Name string `json:"string"`
}

// IsEntity implement gqlgen fedruntime Entity
// func (User) IsEntity() {}
