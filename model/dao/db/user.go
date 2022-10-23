package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserDao interface {
	AddUser()
	UpdateUserByUID()
	DeleteUserByUID()
	FindUserByUID()
	FindUserList()
}

func NewUser() *user {
	return &user{}
}

func NewUserList() *userList {
	return &userList{}
}

func (user *user) AddUser(ctx context.Context) (*user, error) {
	err := GlobalDb.Table(user.TableName()).Create(user).Error
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return user, nil
}

func (user *user) UpdateUserByUID(ctx context.Context, uid uint64, values map[string]interface{}) (int64, error) {
	var rowsAffected int64
	res := GlobalDb.Table(user.TableName()).Where("uid = ?", uid).Updates(values)
	if res.Error != nil {
		log.Println("err", res.Error)
		return rowsAffected, res.Error
	}
	return res.RowsAffected, nil
}

func (user *user) DeleteUserByUID(ctx context.Context, uid uint64) (int64, error) {
	var rowsAffected int64

	res := GlobalDb.Table(user.TableName()).Where("uid = ?", uid).Update(UserColumn.IsDel, 1)
	if res.Error != nil {
		log.Println("err", res.Error)
		return rowsAffected, res.Error
	}
	return res.RowsAffected, nil
}

func (user *user) FindUserByEmail(ctx context.Context, email string) (*user, error) {
	params := map[string]interface{}{
		UserColumn.Email: email,
	}
	res := GlobalDb.Table(user.TableName()).Where(params).Limit(1).Find(user)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return user, nil
}

func (user *user) FindUserList(ctx context.Context, pn int, num int, params map[string]interface{}) (*userList, error) {
	userList := NewUserList()
	offset := pn * num

	res := GlobalDb.Table(user.TableName()).Offset(offset).Limit(num).Where(params).Find(userList)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return userList, nil
}

func (user *user) BeforeCreate(tx *gorm.DB) error {
	if user.CreatedAt.IsZero() {
		t := time.Now()
		user.CreatedAt = &t
	}

	if user.UpdatedAt.IsZero() {
		t := time.Now()
		user.UpdatedAt = &t
	}

	return nil
}

func (user *user) BeforeUpdate(tx *gorm.DB) error {
	if values, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		if _, ok := values[UserColumn.UpdatedAt]; !ok {
			t := time.Now()
			values[UserColumn.UpdatedAt] = &t
		}

		if _, ok := values[UserColumn.Version]; !ok {
			values[UserColumn.UpdatedAt] = gorm.Expr(fmt.Sprintf("%s + ?", UserColumn.Version), 1)
		}
	}

	return nil
}
