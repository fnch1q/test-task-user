package usecase

import (
	"context"
	"test-task-user/internal/entities"
	"test-task-user/internal/infrastructure/enrichment"
	"time"
)

type CreateUserInput struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic"`
}

type CreateUserUseCase struct {
	userRepo   entities.UserRepository
	enrichment enrichment.Client
	ctxTimeout time.Duration
}

func NewCreateUserUseCase(
	userRepo entities.UserRepository,
	enrichment enrichment.Client,
	t time.Duration,
) CreateUserUseCase {
	return CreateUserUseCase{
		userRepo:   userRepo,
		enrichment: enrichment,
		ctxTimeout: t,
	}
}

func (uc CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) error {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	if err := uc.validateInput(input); err != nil {
		return err
	}

	age, err := uc.enrichment.GetAge(input.Name)
	if err != nil {
		return err
	}

	gender, err := uc.enrichment.GetGender(input.Name)
	if err != nil {
		return err
	}

	country, err := uc.enrichment.GetNationality(input.Name)
	if err != nil {
		return err
	}

	user := entities.NewUserCreate(
		input.Name,
		input.Surname,
		age,
		gender,
		country,
		time.Now(),
	)

	if input.Patronymic != nil {
		user.SetPatronymic(input.Patronymic)
	}

	_, err = uc.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (uc CreateUserUseCase) validateInput(input CreateUserInput) error {
	if input.Name == "" {
		return entities.ErrInvalidName
	}
	if input.Surname == "" {
		return entities.ErrInvalidSurname
	}
	return nil
}
