package signup

import (
	"auth/dto"
	"context"
)

type SignController struct {
	useCase SignUseCase
}

func (c SignController) Signup(context context.Context, payload RegisterReq) (dto.BaseResponse, error) {
	return c.useCase.Signup(context, payload)
}
