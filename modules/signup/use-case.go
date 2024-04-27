package signup

import (
	"context"
	"github.com/iman-khaeruddin/isport-auth/dto"
	"github.com/iman-khaeruddin/isport-auth/entity"
	"github.com/iman-khaeruddin/isport-auth/repository"
	"github.com/iman-khaeruddin/isport-auth/utils/hash"
	"time"
)

type SignUseCase struct {
	userRepo repository.UserInterface
}

func NewSignUseCase(userRepo repository.UserInterface) SignUseCase {
	return SignUseCase{userRepo: userRepo}
}

type SignInterface interface {
	Signup(ctx context.Context, payload RegisterReq) (dto.BaseResponse, error)
}

func (uc SignUseCase) Signup(ctx context.Context, payload RegisterReq) (dto.BaseResponse, error) {
	var user *entity.User
	hashPass, _ := hash.CreateSignature([]byte(payload.Password))
	user = &entity.User{
		Email:       payload.Email,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		Password:    hashPass,
		CreatedTime: time.Now(),
		IsDeleted:   false,
	}
	_, err := uc.userRepo.Create(ctx, user)
	if err != nil {
		return dto.DefaultErrorBaseResponse(err), err
	}

	return dto.DefaultSuccessResponseWithMessage("register success"), nil
}
