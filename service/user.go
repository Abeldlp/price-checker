package service

import (
	"fmt"
	"os"
	"strconv"

	"githumb.com/Abeldlp/price-checker/config"
	"githumb.com/Abeldlp/price-checker/model"
	"gopkg.in/gomail.v2"
)

type UserService interface {
	GetUser(userId int) (*model.User, error)
	SaveUser(userEmail string)
	NotifyUser(userEmail string, product model.Product)
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

// 値段が下がり次第登録ユーザにお知らせのメールを送信
func (p *UService) NotifyUser(userEmail string, product model.Product) {

	// メッセージを作成
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_EMAIL"))
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Price Checker Notification")
	m.SetBody("text/plain", fmt.Sprintf("The product you were waiting for is now: %d. Go check it out in: %s", product.CurrentPrice, product.Url))

	// SMTPDialer を作成しGmailに認証
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("GMAIL_EMAIL"), os.Getenv("GMAIL_PASSWORD"))

	// メールを送信
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
