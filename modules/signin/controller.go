package signin

import (
	"context"
	"github.com/iman-khaeruddin/isport-auth/dto"
)

type SignController struct {
	useCase SignUseCase
}

func (c SignController) Signin(context context.Context, payload LoginReq) (dto.BaseResponse, error) {
	return c.useCase.Signin(context, payload)
}
