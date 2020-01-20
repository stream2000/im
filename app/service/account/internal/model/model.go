package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64
	Title   string
	Content string
	Author  string
}

type Account struct {
	ID       string
	Email    string
	Password string
	//Description string
	//PhoneNumber string
	//Birthday xtime.Time
}
