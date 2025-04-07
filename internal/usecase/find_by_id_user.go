package usecase

import (
	"context"
	"test-task-user/internal/entities"
	"time"
)

type FindUserByIDUseCase struct {
	userRepo   entities.UserRepository
	ctxTimeout time.Duration
}

type FindUserByIDInput struct {
	ID uint32
}

type FindUserByIDOutput struct {
	Name        string  `json:"name"`
	Surname     string  `json:"surname"`
	Patronymic  *string `json:"patronymic"`
	Age         uint32  `json:"age"`
	Gender      string  `json:"gender"`
	Nationality string  `json:"nationality"`
}

func NewFindUserByIDUseCase(
	userRepo entities.UserRepository,
	t time.Duration,
) FindUserByIDUseCase {
	return FindUserByIDUseCase{
		userRepo:   userRepo,
		ctxTimeout: t,
	}
}

func (uc FindUserByIDUseCase) Execute(ctx context.Context, input FindUserByIDInput) (FindUserByIDOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	user, err := uc.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		return FindUserByIDOutput{}, err
	}
	return FindUserByIDOutput{
		Name:        user.Name(),
		Surname:     user.Surname(),
		Patronymic:  user.Patronymic(),
		Age:         user.Age(),
		Gender:      user.Gender(),
		Nationality: user.Nationality(),
	}, nil
}
