package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSub(ctx *gin.Context) uint {
	anySub, isExist := ctx.Get("sub")
	if !isExist {
		res := BuildErrorResponse("Failed to edit user", "Sub not exist")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return 0
	}

	sub := anySub.(float64)
	return uint(sub)
}
