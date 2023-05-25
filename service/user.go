package service

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"githumb.com/Abeldlp/price-checker/config"
	"githumb.com/Abeldlp/price-checker/model"
	"gopkg.in/gomail.v2"
)

// UserService ユーザサービス
type UserService interface {
	GetUser(userId int) (*model.User, error)
	SaveUser(userEmail string) int
	NotifyUser(userEmail string, product model.Product)
}

type UService struct{}

// NewUserService ユーザサービスを作成
func NewUserService() UserService {
	return &UService{}
}

// SaveUser ユーザを保存
func (p *UService) SaveUser(userEmail string) int {
	user := model.NewUser(userEmail)

	qry := "INSERT INTO users (email) VALUES ('" + user.Email + "') RETURNING id"

	var id int
	err := config.DB.QueryRow(qry).Scan(&id)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}

	user.Id = id

	return user.Id
}

// GetUser ユーザを取得
func (p *UService) GetUser(userId int) (*model.User, error) {
	qry := "SELECT * FROM users WHERE id=" + strconv.Itoa(userId)
	var user model.User

	// ユーザを取得
	rows, err := config.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// ユーザを設定
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email)
		if err != nil {
			panic(err.Error())
		}
	}

	// エラー処理
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return &user, nil
}

// 値段が下がり次第登録ユーザにお知らせのメールを送信
func (p *UService) NotifyUser(userEmail string, product model.Product) {

	// メールの内容を設定
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("GMAIL_EMAIL"))
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Price Checker Notification")
	m.SetBody("text/plain", fmt.Sprintf("The product you were waiting for is now: %d. Go check it out in: %s", product.CurrentPrice, product.Url))

	// SMTPサーバーの情報を設定
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("GMAIL_EMAIL"), os.Getenv("GMAIL_PASSWORD"))

	// メールを送信
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}

	// メールを送信したら成功メッセージを表示
	fmt.Println("Email sent successfully!")
}
