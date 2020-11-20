/*package entity

import "github.com/gin-gonic/gin"

// Flight entity
type Flight struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	start string `json:"start"`
	end   string `json:"end"`
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
*/