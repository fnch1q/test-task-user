package usecase

import (
	"context"
	"test-task-user/internal/entities"
	"time"
)

type UpdateUserInput struct {
	ID          uint32  `json:"-"`
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	Patronymic  *string `json:"patronymic"`
	Age         *uint32 `json:"age"`
	Gender      *string `json:"gender"`
	Nationality *string `json:"nationality"`
}

type UpdateUserUseCase struct {
	userRepo   entities.UserRepository
	ctxTimeout time.Duration
}

func NewUpdateUserUseCase(
	userRepo entities.UserRepository,
	t time.Duration,
) UpdateUserUseCase {
	return UpdateUserUseCase{
		userRepo:   userRepo,
		ctxTimeout: t,
	}
}

func (uc UpdateUserUseCase) Execute(ctx context.Context, input UpdateUserInput) error {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	user, err := uc.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if input.Name != nil {
		user.SetName(*input.Name)
	}

	if input.Surname != nil {
		user.SetSurname(*input.Surname)
	}

	if input.Patronymic != nil {
		user.SetPatronymic(input.Patronymic)
	}

	if input.Age != nil {
		user.SetAge(*input.Age)
	}

	if input.Gender != nil {
		user.SetGender(*input.Gender)
	}

	if input.Nationality != nil {
		user.SetNationality(*input.Nationality)
	}

	now := time.Now()
	user.SetUpdatedAt(&now)

	err = uc.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
