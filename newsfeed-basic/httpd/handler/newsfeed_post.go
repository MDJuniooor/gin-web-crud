package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"newsfeeder/platform/newsfeed"
)

type newfeedPostRequest struct {
	Title string `json:"title"`
	Post string  `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newfeedPostRequest{}
		c.Bind(&requestBody)
		
		item := newsfeed.Item{
			Title: requestBody.Title,
			Post : requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}
