package authorHandler

import (
	"net/http"

	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/request"
	"github.com/brunobotter/casa-codigo/internal/util"
	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	svc contract.ServiceManager
	cfg *mapping.Config
}

func NewAuthorController(cfg *mapping.Config, svc contract.ServiceManager) *AuthorController {
	return &AuthorController{
		svc: svc,
		cfg: cfg,
	}
}

func (s *AuthorController) SaveNewAuthor(ctx *gin.Context) {
	var request request.NewAuthorRequest
	if err := ctx.Bind(&request); err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Invalid Body")
		return
	}
	response, err := s.svc.InternalService().AuthorService().Save(ctx, request)
	if err != nil {
		util.ResponderApiError(ctx, http.StatusBadRequest, err, "Error to save author")
		return
	}
	util.ResponderApiOk(ctx, response)

}
