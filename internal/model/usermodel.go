package model

import (
	"errors"
	"integra_backend/internal/db"
	"integra_backend/internal/entity"

	"github.com/Masterminds/squirrel"
	"github.com/labstack/gommon/log"
)

type userModel struct {
	dbConn db.DbConnection
}

type UserModel interface {
	CreateUser(*entity.UserEntity) (*entity.UserEntity, error)
	ListUsers() ([]*entity.UserEntity, error)
	UpdateUser(*entity.UserEntity, int64) (*entity.UserEntity, error)
	DeleteUser(user_id int64) (*entity.UserEntity, error)
}

func NewUserModel(dbConn db.DbConnection) UserModel {
	return &userModel{
		dbConn: dbConn,
	}
}

func (um *userModel) CreateUser(user *entity.UserEntity) (*entity.UserEntity, error) {
	db := um.dbConn.GetConnection()
	users, err := um.getUserByUserName(user.UserName)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		err = errors.New("userName in use")
		return nil, err
	}
	users, err = um.getUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		err = errors.New("email in use")
		return nil, err
	}
	stmt, err := db.Prepare("INSERT INTO users (user_name,first_name,last_name,email,user_status,department) VALUES ($1,$2,$3,$4,$5,$6) RETURNING user_id")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(user.UserName, user.FirstName, user.LastName, user.Email, user.UserStatus, user.Department).Scan(&user.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (um *userModel) ListUsers() ([]*entity.UserEntity, error) {
	selectBuilder := squirrel.Select("*").From("users").Suffix("ORDER BY user_name")
	users, err := um.scanRows(selectBuilder)
	return users, err
}

func (um *userModel) UpdateUser(user *entity.UserEntity, user_id int64) (*entity.UserEntity, error) {
	db := um.dbConn.GetConnection()
	users, err := um.getUserByUserName(user.UserName)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 && users[0].UserId != user_id {
		err = errors.New("userName in use")
		return nil, err
	}
	users, err = um.getUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 && users[0].UserId != user_id {
		err = errors.New("email in use")
		return nil, err
	}
	stmt, err := db.Prepare("UPDATE users SET user_name = $1, first_name = $2, last_name = $3, email = $4, user_status = $5, department = $6 WHERE user_id = $7 RETURNING user_id")

	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(user.UserName, user.FirstName, user.LastName, user.Email, user.UserStatus, user.Department, user_id).Scan(&user.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (um *userModel) DeleteUser(user_id int64) (*entity.UserEntity, error) {
	db := um.dbConn.GetConnection()
	users, err := um.getUserById(user_id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		err = errors.New("user does not exist")
		return nil, err
	}
	deleteBuilder := squirrel.Delete("users").Where("user_id = $1", user_id)
	_, err = deleteBuilder.RunWith(db).Exec()
	if err != nil {
		return nil, err
	}
	return users[0], nil
}

//Private Functions

func (um *userModel) getUserByUserName(userName string) ([]*entity.UserEntity, error) {
	selectBuilder := squirrel.Select("*").From("users").Where("user_name IN ($1)", userName)
	users, err := um.scanRows(selectBuilder)
	return users, err
}

func (um *userModel) getUserByEmail(email string) ([]*entity.UserEntity, error) {
	selectBuilder := squirrel.Select("*").From("users").Where("email IN ($1)", email)
	users, err := um.scanRows(selectBuilder)
	return users, err
}

func (um *userModel) getUserById(user_id int64) ([]*entity.UserEntity, error) {
	selectBuilder := squirrel.Select("*").From("users").Where("user_id = $1", user_id)
	users, err := um.scanRows(selectBuilder)
	return users, err
}

func (um *userModel) scanRows(selectBuilder squirrel.SelectBuilder) ([]*entity.UserEntity, error) {
	db := um.dbConn.GetConnection()
	rows, err := selectBuilder.RunWith(db).Query()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var users []*entity.UserEntity
	for rows.Next() {
		var user entity.UserEntity
		rows.Scan(&user.UserId, &user.UserName, &user.FirstName, &user.LastName, &user.Email, &user.UserStatus, &user.Department)
		users = append(users, &user)
	}
	return users, nil
}
