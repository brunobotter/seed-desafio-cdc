package util

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAndValidateIntParam(ctx *gin.Context, param string, errorMessage string, isZeroValid bool) (id int64, err error) {
	IDstr := ctx.Param(param)

	if strings.TrimSpace(IDstr) == "" {
		//		err = errors.NewValidationError(param, errorMessage)
		return
	}
	id, err = strconv.ParseInt(IDstr, 10, 64)
	if !isZeroValid && id == 0 {
		//		err = errors.NewValidationError(param, errorMessage)

	}

	return
}
