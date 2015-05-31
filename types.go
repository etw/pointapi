package pointapi

import (
	"time"
)

type PostList struct {
	HasNext bool       `json:"has_next"`
	Posts   []PostMeta `json:"posts"`
}

type PostMeta struct {
	Bookmarked  bool     `json:"bookmarked"`
	Uid         int      `json:"uid"`
	Subscribed  bool     `json:"subscribed"`
	Editable    bool     `json:"editable"`
	Recommended bool     `json:"recommended"`
	Post        PostData `json:"post"`
}

type PostData struct {
	Tags          []string   `json:"tags"`
	CommentsCount int        `json:"comments_count"`
	Author        PostAuthor `json:"author"`
	Text          string     `json:"text"`
	Created       time.Time  `json:"created"`
	Type          string     `json:"type"`
	Id            string     `json:"id"`
	Private       bool       `json:"private"`
}

type PostAuthor struct {
	Login  string `json:"login"`
	Id     int    `json:"id"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}
