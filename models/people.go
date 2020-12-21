package models

import (
	"log"
	"time"
)

type People struct {
	ID int 
	UUID string
	Firstname string
	Lastname string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Peopleの作成
func (p *People) CreatePeople() (err error) {
	cmd := `insert into people (uuid, firstname, lastname, email, 
		password, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?)`
	
	_, err = Db.Exec(cmd, createUUID(), p.Firstname, p.Lastname, p.Email,
		Encrypt(p.Password), time.Now(), time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err 
} 