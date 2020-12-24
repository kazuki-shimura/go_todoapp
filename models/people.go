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
	cmd := `insert into people (
		uuid, firstname, lastname, email, 
		password, created_at, updated_at) 
		values (?, ?, ?, ?, ?, ?, ?)`
	
	_, err = Db.Exec(cmd, createUUID(), p.Firstname, p.Lastname, p.Email,
		Encrypt(p.Password), time.Now(), time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err 
}

// Peopleの取得
func GetPeople(id int) (people People, err error) {
	people = People{}
	cmd := `select id, uuid, firstname, lastname, email, password, 
		created_at, updated_at from people where id = ?`
	
	err = Db.QueryRow(cmd, id).Scan(
		&people.ID,
		&people.UUID,
		&people.Firstname,
		&people.Lastname,
		&people.Email,
		&people.Password,
		&people.CreatedAt,
		&people.UpdatedAt,
	)
	return people, err
}

// // People全取得
// func GetAllPeople() (people People, err error) {
// 	peoples = 
// }

// Peopleの更新
func (p *People) UpdatePeople() (err error) {
	cmd := `update people set firstname = ?, lastname = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, p.Firstname, p.Lastname, p.Email, p.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err 
}

//Peopleの削除
func (p *People) DeletePeople() (err error) {
	cmd := `delete from people where id = ?`
	_, err = Db.Exec(cmd, p.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err 
}