package signin

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/isport-auth/dto"
	"github.com/iman-khaeruddin/isport-auth/repository"
	"github.com/iman-khaeruddin/isport-auth/utils/validator"
	"gorm.io/gorm"
	"net/http"
	"os"
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

func (h SignRequestHandler) HandleSignin(router *gin.Engine) {
	signupUseCase := SignUseCase{
		userRepo: repository.NewUser(h.iSportDB),
	}
	h.ctrl = SignController{useCase: signupUseCase}

	router.POST("/signin", gin.BasicAuth(gin.Accounts{os.Getenv("USERNAME"): os.Getenv("PASSWORD")}), h.signin)
}

func (h SignRequestHandler) signin(c *gin.Context) {
	var request LoginReq
	if !validator.BindAndValidateWithAbort(c, &request) {
		return
	}

	response, err := h.ctrl.Signin(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorBaseResponse(err))
		return
	}
	c.JSON(http.StatusOK, response)
}
