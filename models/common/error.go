package common

type Error struct {
	Code     int         `json:"code"`
	Messages interface{} `json:"messages"`
}