package util

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func CpfCnpjValidator(fl validator.FieldLevel) bool {
	doc := fl.Field().String()
	cpfRegex := `^\d{3}\.\d{3}\.\d{3}-\d{2}$`
	cnpjRegex := `^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`

	matchCPF, _ := regexp.MatchString(cpfRegex, doc)
	matchCNPJ, _ := regexp.MatchString(cnpjRegex, doc)

	return matchCPF || matchCNPJ
}
