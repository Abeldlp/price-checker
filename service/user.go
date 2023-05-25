package service

import (
	"fmt"
	"strconv"

	"githumb.com/Abeldlp/price-checker/config"
	"githumb.com/Abeldlp/price-checker/model"
)

type UserService interface {
	GetUser(userId int) (*model.User, error)
	SaveUser(userEmail string)
	NotifyUser(userEmail string)
}

type UService struct{}

func NewUserService() UserService {
	return &UService{}
}

func (p *UService) SaveUser(userEmail string) {
	user := model.NewUser(userEmail)

	fmt.Println("Creating user", user)
}

func (p *UService) GetUser(userId int) (*model.User, error) {
	qry := "SELECT * FROM users WHERE id=" + strconv.Itoa(userId)
	var user model.User

	rows, err := config.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email)
		if err != nil {
			panic(err.Error())
		}
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return &user, nil
}

func (p *UService) NotifyUser(userEmail string) {
	fmt.Println("Sending message to:", userEmail)
}
