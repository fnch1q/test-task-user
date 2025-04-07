package usecase

import (
	"context"
	"test-task-user/internal/entities"
	"time"
)

type DeleteUserInput struct {
	ID uint32
}

type DeleteUserUseCase struct {
	userRepo   entities.UserRepository
	ctxTimeout time.Duration
}

func NewDeleteUserUseCase(
	userRepo entities.UserRepository,
	t time.Duration,
) DeleteUserUseCase {
	return DeleteUserUseCase{
		userRepo:   userRepo,
		ctxTimeout: t,
	}
}

func (uc DeleteUserUseCase) Execute(ctx context.Context, input DeleteUserInput) error {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	if err := uc.userRepo.Delete(ctx, input.ID); err != nil {
		return err
	}

	return nil
}
