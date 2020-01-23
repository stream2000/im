package model

// cache model
type Group struct {
	Name        string
	Id          int64
	Description string
	Members     []int64
}

type GroupMembers struct {
	GroupId   int64
	MembersId []int64
}
