package repo

import (
	"context"
	"test-task-user/internal/entities"
	"test-task-user/internal/infrastructure/database"
	"time"

	"gorm.io/gorm"
)

type userGORM struct {
	ID          uint32     `gorm:"primary_key"`
	Name        string     `gorm:"column:name"`
	Surname     string     `gorm:"column:surname"`
	Patronymic  *string    `gorm:"column:patronymic"`
	Age         uint32     `gorm:"column:age"`
	Gender      string     `gorm:"column:gender"`
	Nationality string     `gorm:"column:nationality"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}

type UserDB struct {
	dbManager database.DBManager
	tableName string
}

func NewUserDB(dbManager database.DBManager) UserDB {
	return UserDB{
		dbManager: dbManager,
		tableName: "users",
	}
}

func (u UserDB) Create(ctx context.Context, user entities.User) (entities.User, error) {
	var userGORM = userGORM{
		Name:        user.Name(),
		Surname:     user.Surname(),
		Patronymic:  user.Patronymic(),
		Age:         user.Age(),
		Gender:      user.Gender(),
		Nationality: user.Nationality(),
	}

	if err := u.dbManager.With(ctx).Table(u.tableName).Create(&userGORM).Error; err != nil {
		return entities.User{}, err
	}

	user.SetID(userGORM.ID)
	return user, nil
}

func (u UserDB) Update(ctx context.Context, user entities.User) error {
	updatesMap := map[string]interface{}{
		"name":        user.Name(),
		"surname":     user.Surname(),
		"patronymic":  user.Patronymic(),
		"age":         user.Age(),
		"gender":      user.Gender(),
		"nationality": user.Nationality(),
		"created_at":  user.CreatedAt(),
		"updated_at":  user.UpdatedAt(),
	}

	if err := u.dbManager.With(ctx).Table(u.tableName).Where("id = ?", user.ID()).Updates(updatesMap).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return entities.ErrUserNotFound
		default:
			return err
		}
	}
	return nil
}

func (u UserDB) Delete(ctx context.Context, id uint32) error {
	if err := u.dbManager.With(ctx).Table(u.tableName).Where("id = ?", id).Delete(&userGORM{}).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return entities.ErrUserNotFound
		default:
			return err
		}
	}
	return nil
}

func (u UserDB) FindByID(ctx context.Context, id uint32) (entities.User, error) {
	var userGORM userGORM

	if err := u.dbManager.With(ctx).Table(u.tableName).Where("id = ?", id).First(&userGORM).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return entities.User{}, entities.ErrUserNotFound
		}
		return entities.User{}, err
	}

	return entities.NewUser(
		userGORM.ID,
		userGORM.Name,
		userGORM.Surname,
		userGORM.Patronymic,
		userGORM.Age,
		userGORM.Gender,
		userGORM.Nationality,
		userGORM.CreatedAt,
		userGORM.UpdatedAt,
	), nil
}

func (u UserDB) FindAll(ctx context.Context, searchParams entities.UserSearchParams, page, limit int64) ([]entities.User, int64, error) {
	return nil, 0, nil
}
