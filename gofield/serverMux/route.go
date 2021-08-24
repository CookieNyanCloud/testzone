package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init()  {
	posts = []Post{
		{
			Id:    1,
			Title: "Title 1",
			Text:  "Text 1",
		},
	}
}

func GetPosts(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-type", "application/json")
	result, err:= json.Marshal(posts)
	if err!=nil{
		resp.WriteHeader(http.StatusInternalServerError)
		_,_=resp.Write([]byte(`{"error":"Error marshaling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	_,_=resp.Write(result)
}