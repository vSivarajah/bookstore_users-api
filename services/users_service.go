package services

import (
	"github.com/vSivarajah/bookstore_users-api/domain/users"
	"github.com/vSivarajah/bookstore_users-api/utils/date_utils"
	"github.com/vSivarajah/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowString()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

//CREATE TABLE `users_db`.`users`(`id` BIGINT(20) NOT NULL AUTO_INCREMENT, `firstname` VARCHAR(45), `lastname` VARCHAR(45), `email` VARCHAR(45), `datecreated` VARCHAR(45), PRIMARY KEY (`id`), UNIQUE INDEX `email_UNIQUE`(`email` ASC));
