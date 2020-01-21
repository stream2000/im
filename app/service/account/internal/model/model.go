package model

type Article struct {
	ID      int64
	Title   string
	Content string
	Author  string
}

type Account struct {
	UID           string
	Email         string
	Password      string
	NickName      string
	Sign          string
	ProfilePicUrl string
}
