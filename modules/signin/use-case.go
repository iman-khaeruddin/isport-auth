package signin

import (
	"context"
	"github.com/iman-khaeruddin/isport-auth/dto"
	"github.com/iman-khaeruddin/isport-auth/entity"
	"github.com/iman-khaeruddin/isport-auth/repository"
	"github.com/iman-khaeruddin/isport-auth/utils/hash"
	"gorm.io/gorm"
	"time"
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

	updateLastLogin := &entity.User{
		ID:        user.ID,
		LastLogin: time.Now(),
	}
	uc.userRepo.UpdateSelectedFields(ctx, updateLastLogin, "LastLogin")

	return dto.BaseResponse{
		Data:         token,
		Success:      true,
		MessageTitle: "",
		Message:      "login success",
	}, nil
}
