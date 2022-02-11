package controller

import (
	".../entity/registration.go"
	"github.com/gin-gonic/gin"
)

func Save(ctx *gin.Context) error {
	var req registration.RegistrationRequest
	err := ctx.shouldBindJSON(&registration)
	if err != nil {
		return err
	}
	return nil
}
