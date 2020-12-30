package models

import (
	"log"
	"time"
)

type AAA struct {
	Id int
	Uuid string
	First string
	Second string
	Third string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// aaaの作成
func (a *AAA) CreateAAA() (err error) {
	cmd := `insert into aaa (
		uuid, first, second, third, email, password,
		created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?)`
	
	_, err = Db.Exec(cmd, createUUID(), a.First, a.Second, a.Third,
		a.Email, Encrypt(a.Password), time.Now(), time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}