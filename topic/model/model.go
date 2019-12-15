package model

type Topic struct {
	Uid        int32  `json:"uid"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Like       int32  `json:"like"`
	CreateTime string `json:"create_time"`
}

type TopicJoinUser struct {
	Uid        int32  `json:"uid"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Like       int32  `json:"like"`
	CreateTime string `json:"create_time"`
	Name  string
}

type Comment struct {
	Content string `json:"content"`
	Uid     int    `json:"uid"`
}
