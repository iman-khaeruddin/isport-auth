package signin

import (
	"auth/dto"
	"context"
)

type SignController struct {
	useCase SignUseCase
}

func (c SignController) Signin(context context.Context, payload LoginReq) (dto.BaseResponse, error) {
	return c.useCase.Signin(context, payload)
}
