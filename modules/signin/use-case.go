package signin

import (
	"auth/dto"
	"auth/repository"
	"auth/utils/hash"
	"context"
	"gorm.io/gorm"
)

type SignUseCase struct {
	userRepo repository.UserInterface
}

func NewSignUseCase(userRepo repository.UserInterface) SignUseCase {
	return SignUseCase{userRepo: userRepo}
}

type SignInterface interface {
	Signin(ctx context.Context, payload LoginReq) (dto.BaseResponse, error)
}

func (uc SignUseCase) Signin(ctx context.Context, payload LoginReq) (dto.BaseResponse, error) {
	user, err := uc.userRepo.FindByEmail(ctx, payload.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dto.DefaultErrorBaseResponseWithMessage("email tidak dikenal"), nil
		}
		return dto.DefaultErrorBaseResponse(err), err
	}

	r, err := hash.Validate([]byte(payload.Password), user.Password)
	if !r || err != nil {
		return dto.DefaultErrorBaseResponseWithMessage("email atau password salah"), nil
	}

	token, err := hash.CreateToken(user)
	if err != nil {
		return dto.DefaultErrorBaseResponse(err), err
	}

	return dto.BaseResponse{
		Data:         token,
		Success:      true,
		MessageTitle: "",
		Message:      "login success",
	}, nil
}
