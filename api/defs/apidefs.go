package defs

// requests

type UserCredential struct {
	Username string `json:"username"`
	Pwd string `json:"pwd"`
}

// Response
type SignedUp struct {
	Success bool `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

type SimpleSession struct {
	Username string // login name
	TTL int64
}