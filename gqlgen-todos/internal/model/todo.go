package model

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user_id"`
}

// IsEntity implement gqlgen fedruntime Entity
// func (Todo) IsEntity() {}
