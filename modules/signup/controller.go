package signup

import (
	"context"
	"github.com/iman-khaeruddin/isport-auth/dto"
)

type SignController struct {
	useCase SignUseCase
}

func (c SignController) Signup(context context.Context, payload RegisterReq) (dto.BaseResponse, error) {
	return c.useCase.Signup(context, payload)
}
