/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 17:02:54
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-23 21:47:37
 */
package svruser

import (
	"blogserver/model/dao/db"
	"context"
	"fmt"
	"log"
	"time"
)

type UserService interface {
	AddUser()
	UpdateUserByUID()
	DeleteUserByUID()
	FindUserByEmail()
	FindUserList()
}

type User struct {
	ID        uint64     `json:"id"`
	UID       uint64     `json:"uid"`
	Password  string     `json:"password"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"create_at"`
	UpdatedAt *time.Time `json:"update_at"`
}

type UserList []User

type UpdateUserParams struct {
	Password *string
	Username *string
}

func NewUserService() *User {
	return &User{}
}

func (user *User) AddUser(ctx context.Context) (*User, error) {
	dbUser := db.NewUser()

	dbUser.UID = user.UID
	dbUser.Password = user.Password
	dbUser.Username = user.Username
	dbUser.Email = user.Email

	dbResult, err := dbUser.AddUser(ctx)

	if err != nil {
		msg := fmt.Errorf("SERVICE_USER_USER_AddUser_Failed")
		log.Println("err", msg)
		return nil, msg
	}

	res := &User{
		ID:        dbResult.ID,
		UID:       dbResult.UID,
		Username:  dbResult.Username,
		Password:  dbResult.Password,
		Email:     dbResult.Email,
		CreatedAt: dbResult.CreatedAt,
		UpdatedAt: dbResult.UpdatedAt,
	}

	return res, nil
}

func (user *User) UpdateUserByUID(ctx context.Context, uid uint64, params *UpdateUserParams) error {
	dbUser := db.NewUser()

	val := map[string]interface{}{}
	if params.Username != nil {
		val[db.UserColumn.Username] = *params.Username
	}

	if params.Password != nil {
		val[db.UserColumn.Password] = *params.Password
	}

	_, err := dbUser.UpdateUserByUID(ctx, uid, val)

	if err != nil {
		msg := fmt.Errorf("SERVICE_USER_USER_UpdateUserByUID_DbUserUpdateUserByUID_Failed")
		log.Println("err", msg)
		return msg
	}

	return nil
}

func (user *User) DeleteUserByUID(ctx context.Context, uid uint64) error {
	dbUser := db.NewUser()

	_, err := dbUser.DeleteUserByUID(ctx, uid)

	if err != nil {
		msg := fmt.Errorf("SERVICE_USER_USER_DeleteUserByUID_DbUserDeleteUserByUID_Failed")
		log.Println("err", msg)
		return msg
	}

	return nil
}

func (user *User) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	dbUser := db.NewUser()

	dbResult, err := dbUser.FindUserByEmail(ctx, email)

	if err != nil {
		msg := fmt.Errorf("SERVICE_USER_USER_FindUserByEmail_DbUserFindUserByEmail_Failed")
		log.Println("err", msg)
		return nil, msg
	}

	result := &User{
		ID:        dbResult.ID,
		UID:       dbResult.UID,
		Username:  dbResult.Username,
		Email:     dbResult.Email,
		Password:  dbResult.Password,
		CreatedAt: dbResult.CreatedAt,
		UpdatedAt: dbResult.UpdatedAt,
	}

	return result, nil

}

func (user *User) FindUserList(ctx context.Context, pn int, num int, params map[string]interface{}) (*UserList, error) {
	dbUser := db.NewUser()

	dbResult, err := dbUser.FindUserList(ctx, pn, num, params)

	if err != nil {
		msg := fmt.Errorf("SERVICE_USER_USER_FindUserList_DbUserFindUserList_Failed")
		log.Println("err", msg)
		return nil, msg
	}

	result := make(UserList, 0)

	for _, item := range *dbResult {
		temp := User{
			ID:       item.ID,
			UID:      item.UID,
			Username: item.Username,
			Password: item.Password,
			Email:    item.Email,
		}

		result = append(result, temp)
	}

	return &result, nil

}
