package repository

import (
	"belajar-go-echo/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllUser(t *testing.T) {
	tu := NewTestUnit()

	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "email", "password"}).
				AddRow(1, "ristiaamandari@gmail.com").AddRow(2, "123def"))

	listUser, err := tu.IFaceUserRepo.GetAllUser()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(users)
}

func TestCreateUser(t *testing.T) {
	tu := NewTestUnit()
	user := model.User{
		Id:       1,
		Username: "ristiaamandari@mail.com",
		Password: "123edf",
	}

	tu.Mock.ExpectBegin()
	tu.Mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`,`id`) VALUES (?,?,?)")).
		WithArgs(user.Email, user.Password, user.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	tu.Mock.ExpectCommit()
	err := tu.IFaceUserRepo.CreateUser(user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success insert")
}
