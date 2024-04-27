package signup

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/isport-auth/dto"
	"github.com/iman-khaeruddin/isport-auth/repository"
	"github.com/iman-khaeruddin/isport-auth/utils/validator"
	"gorm.io/gorm"
	"net/http"
)

type SignRequestHandler struct {
	iSportDB *gorm.DB
	ctrl     SignController
}

func NewSignRequestHandler(iSportDB *gorm.DB) SignRequestHandler {
	return SignRequestHandler{
		iSportDB: iSportDB,
	}
}

func (h SignRequestHandler) HandleSignup(router *gin.Engine) {
	signupUseCase := SignUseCase{
		userRepo: repository.NewUser(h.iSportDB),
	}
	h.ctrl = SignController{useCase: signupUseCase}

	router.POST("/signup", h.signup)
}

func (h SignRequestHandler) signup(c *gin.Context) {
	var request RegisterReq
	if !validator.BindAndValidateWithAbort(c, &request) {
		return
	}

	response, err := h.ctrl.Signup(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorBaseResponse(err))
		return
	}
	c.JSON(http.StatusOK, response)
}
