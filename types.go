package pointapi

const POINTTIMELAYOUT = "2006-01-02T15:04:05.999999"

// Struct with list of posts objects
type PostList struct {
	HasNext bool       `json:"has_next"`
	Posts   []PostMeta `json:"posts"`
}

// Struct with post metadata
type PostMeta struct {
	Bookmarked  bool     `json:"bookmarked"`
	Uid         int      `json:"uid"`
	Subscribed  bool     `json:"subscribed"`
	Editable    bool     `json:"editable"`
	Recommended bool     `json:"recommended"`
	Post        PostData `json:"post"`
}

// Struct with post data
type PostData struct {
	Tags          []string   `json:"tags"`
	CommentsCount int        `json:"comments_count"`
	Author        PostAuthor `json:"author"`
	Created       string     `json:"created"`
	Type          string     `json:"type"`
	Id            string     `json:"id"`
	Private       bool       `json:"private"`
	Files         []string   `json:"files"`
	Text          string     `json:"text"`
}

// Struct with info about user
type PostAuthor struct {
	Login  string `json:"login"`
	Id     int    `json:"id"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}
