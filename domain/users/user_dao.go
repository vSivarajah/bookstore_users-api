package users

import (
	"fmt"
	"strings"

	"github.com/vSivarajah/bookstore_users-api/datasources/mysql/users_db"
	"github.com/vSivarajah/bookstore_users-api/utils/errors"
)

const (
	errorNoRows           = "no rows in result set"
	indexUniqueEmail      = "unique_EMAIL"
	queryInsertUser       = "INSERT INTO users(firstname, lastname, email,status, password, datecreated) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, firstname,lastname, email, status, datecreated FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET firstname=?, lastname=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT ID, firstname, lastname, email, status, datecreated FROM users WHERE status=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d note found", user.ID))
		}
		fmt.Println(err)
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.ID, err.Error()))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", err.Error()))
	}
	user.ID = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to update user: %s", err.Error()))
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	if _, err = stmt.Exec(user.ID); err != nil {
		return errors.NewBadRequestError("Cant delete user")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	defer rows.Close()

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError("parse error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
	}
	return results, nil
}
