import (
	"belajar-go-echo/model"
	"fmt"

	"gorm.io/gorm"
)

type IUserRepo interface {
	GetAllUser() ([]model.UserCostum, error)
	CreateUser(user model.User) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &UserRepo{db}
}

func (r UserRepo) GetAllUser() ([]model.UserCostum, error) {
	users := []model.UserCostum{}
	err := r.db.Find(&users).Error // SELECT * FROM users;
	if err != nil {
		fmt.Println("error while GetAllUser", err)
	}
	return users, err
}

func (repo UserRepo) CreateUser(user model.User) error {
	err := repo.db.Create(&user).Error
	if err != nil {
		fmt.Println("error while CreateUser", err)
	}
	return err
}

