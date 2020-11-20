package entity

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Flight entity
type Flight struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Start    string `json:"start"` //yyyy-MM-dd HH:mm
	End      string `json:"end"`   //yyyy-MM-dd HH:mm
	Aircraft string `json:"aircraft"`
	IdAgency int64  `json:"idAgency"`
}

func (a Flight) IsValid() bool {
	valid1 := isValidDate(a.Start)
	valid2 := isValidDate(a.End)

	return a.Name != "" && a.Start != "" && a.End != "" && valid1 && valid2
}

func (a Flight) ToJson() gin.H {
	return gin.H{
		"id":       a.ID,
		"name":     a.Name,
		"start":    a.Start,
		"end":      a.End,
		"aircraft": a.Aircraft,
	}
}

func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02 15:04", date)
	return err == nil
}
