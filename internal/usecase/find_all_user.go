package usecase

import (
	"context"
	"test-task-user/internal/entities"
	"time"
)

type FindAllUsersUseCase struct {
	userRepo   entities.UserRepository
	ctxTimeout time.Duration
}

type FindAllUsersInput struct {
	Name        *string `form:"name"`
	Surname     *string `form:"surname"`
	Patronymic  *string `form:"patronymic"`
	Age         *uint32 `form:"age"`
	Gender      *string `form:"gender"`
	Nationality *string `form:"nationality"`
	Page        int     `form:"page"`
	Limit       int     `form:"limit"`
}

type FindAllUsersOutput struct {
	Users []struct {
		Name        string  `json:"name"`
		Surname     string  `json:"surname"`
		Patronymic  *string `json:"patronymic"`
		Age         uint32  `json:"age"`
		Gender      string  `json:"gender"`
		Nationality string  `json:"nationality"`
	} `json:"users"`
	TotalCount int64 `json:"total_count"`
}

func NewFindAllUsersUseCase(
	userRepo entities.UserRepository,
	t time.Duration,
) FindAllUsersUseCase {
	return FindAllUsersUseCase{
		userRepo:   userRepo,
		ctxTimeout: t,
	}
}

func (uc FindAllUsersUseCase) Execute(ctx context.Context, input FindAllUsersInput) (FindAllUsersOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	users, totalCount, err := uc.userRepo.FindAll(ctx, entities.UserSearchParams{
		Name:        input.Name,
		Surname:     input.Surname,
		Patronymic:  input.Patronymic,
		Age:         input.Age,
		Gender:      input.Gender,
		Nationality: input.Nationality,
	}, input.Page, input.Limit)
	if err != nil {
		return FindAllUsersOutput{}, err
	}

	var usersOutput []struct {
		Name        string  `json:"name"`
		Surname     string  `json:"surname"`
		Patronymic  *string `json:"patronymic"`
		Age         uint32  `json:"age"`
		Gender      string  `json:"gender"`
		Nationality string  `json:"nationality"`
	}
	for _, user := range users {
		usersOutput = append(usersOutput, struct {
			Name        string  `json:"name"`
			Surname     string  `json:"surname"`
			Patronymic  *string `json:"patronymic"`
			Age         uint32  `json:"age"`
			Gender      string  `json:"gender"`
			Nationality string  `json:"nationality"`
		}{
			Name:        user.Name(),
			Surname:     user.Surname(),
			Patronymic:  user.Patronymic(),
			Age:         user.Age(),
			Gender:      user.Gender(),
			Nationality: user.Nationality(),
		})
	}

	return FindAllUsersOutput{
		Users:      usersOutput,
		TotalCount: totalCount,
	}, nil
}
