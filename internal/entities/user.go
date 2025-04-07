package entities

import (
	"context"
	"net/http"
	"test-task-user/internal/errors"
	"time"
)

type User struct {
	id          uint32
	name        string
	surname     string
	patronymic  *string
	age         uint32
	gender      string
	nationality string
	createdAt   time.Time
	updatedAt   *time.Time
}

type UserSearchParams struct {
	Name        *string
	Surname     *string
	Patronymic  *string
	Age         *uint32
	Gender      *string
	Nationality *string
}

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id uint32) error
	FindByID(ctx context.Context, id uint32) (User, error)
	FindAll(ctx context.Context, searchParams UserSearchParams, page, limit int) ([]User, int64, error)
}

var (
	ErrUserNotFound   = errors.NewError(http.StatusNotFound, "user_not_found")
	ErrInvalidName    = errors.NewError(http.StatusBadRequest, "invalid_name")
	ErrInvalidSurname = errors.NewError(http.StatusBadRequest, "invalid_surname")
)

func NewUser(
	id uint32,
	name string,
	surname string,
	patronymic *string,
	age uint32,
	gender string,
	nationality string,
	createdAt time.Time,
	updatedAt *time.Time,
) User {
	return User{
		id:          id,
		name:        name,
		surname:     surname,
		patronymic:  patronymic,
		age:         age,
		gender:      gender,
		nationality: nationality,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func NewUserCreate(
	name string,
	surname string,
	age uint32,
	gender string,
	nationality string,
	createdAt time.Time,
) User {
	return User{
		name:        name,
		surname:     surname,
		age:         age,
		gender:      gender,
		nationality: nationality,
		createdAt:   createdAt,
	}
}

func (u User) ID() uint32 {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Surname() string {
	return u.surname
}

func (u User) Patronymic() *string {
	return u.patronymic
}

func (u User) Age() uint32 {
	return u.age
}

func (u User) Gender() string {
	return u.gender
}

func (u User) Nationality() string {
	return u.nationality
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() *time.Time {
	return u.updatedAt
}

func (u *User) SetID(id uint32) {
	u.id = id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetSurname(surname string) {
	u.surname = surname
}

func (u *User) SetPatronymic(patronymic *string) {
	u.patronymic = patronymic
}

func (u *User) SetAge(age uint32) {
	u.age = age
}

func (u *User) SetGender(gender string) {
	u.gender = gender
}

func (u *User) SetNationality(nationality string) {
	u.nationality = nationality
}

func (u *User) SetUpdatedAt(updatedAt *time.Time) {
	u.updatedAt = updatedAt
}
