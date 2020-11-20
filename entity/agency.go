package entity

import "github.com/gin-gonic/gin"

// Agency entity
type Agency struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (a Agency) IsValid() bool {
	return a.Name != ""
}

func (a Agency) ToJson() gin.H {
	return gin.H{
		"id":   a.ID,
		"name": a.Name,
	}
}
