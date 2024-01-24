package domain

type CreateUser struct {
	Name string `json:"name"`
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
