package repository

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/sawitpro/technical_test/entity"
	"gorm.io/gorm"
)

const charset = `abcdefghijklmnopqrstuvwxyz1234567890`

func (r *repositoryCtx) Now() time.Time {
	return time.Now()
}

func (r *repositoryCtx) RandomString(length int) string {
	value := make([]byte, length)
	for i := range value {
		value[i] = charset[rand.Intn(len(charset))]
	}
	return string(value)
}

func (r *repositoryCtx) GetUserByID(ctx context.Context, userID int) (*entity.User, error) {
	var (
		user = &entity.User{}
		err  error
	)

	db := r.cfg.DB.WithContext(ctx)

	err = db.First(user, "id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (r *repositoryCtx) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	var (
		user = &entity.User{}
		err  error
	)

	db := r.cfg.DB.WithContext(ctx)

	err = db.First(user, "phone_number = ?", phoneNumber).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (r *repositoryCtx) IncrementSuccessfulLogin(ctx context.Context, userID int) error {
	var (
		err error
	)

	db := r.cfg.DB.WithContext(ctx)

	err = db.Model(&entity.User{}).Where(`id = ?`, userID).Update("successful_login", gorm.Expr("successful_login + ?", 1)).Error
	if err != nil {
		log.Printf(`Increment successful login error %s`, err.Error())
		return err
	}

	return nil
}

func (r *repositoryCtx) UpdateProfile(ctx context.Context, user *entity.User) error {
	var (
		err error
	)

	db := r.cfg.DB.WithContext(ctx)

	data := &entity.User{
		PhoneNumber: user.PhoneNumber,
		FullName:    user.FullName,
		UpdatedAt:   user.UpdatedAt,
	}

	err = db.Model(user).Updates(data).Error
	if err != nil {
		log.Printf(`Update profile error %s`, err.Error())
		return err
	}

	return nil
}

func (r *repositoryCtx) Create(ctx context.Context, user *entity.User) error {
	var (
		err error
	)

	db := r.cfg.DB.WithContext(ctx)

	err = db.Create(user).Error
	if err != nil {
		log.Printf(`Create user Error %s`, err.Error())
		return err
	}

	return nil
}
