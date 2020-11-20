package entity

import "github.com/gin-gonic/gin"

type Entity interface {
	IsValid() bool
	ToJson() gin.H
}
