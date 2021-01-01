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

// aaaの取得
func GetAAA(id int) (aaa AAA, err error) {
	aaa = AAA{}
	cmd := `select id, uuid, first, second, third, email, password, created_at, updated_at
		from aaa where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&aaa.Id,
		&aaa.Uuid,
		&aaa.First,
		&aaa.Second,
		&aaa.Third,
		&aaa.Email,
		&aaa.Password,
		&aaa.CreatedAt,
		&aaa.UpdatedAt,
	)
	return aaa, err
}

// aaaの全取得
func GetAAAs() (aaas []AAA, err error) {
	cmd := `select id, uuid, first, second, third, email, password,
	created_at, updated_at from aaa`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var aaa AAA
		err = rows.Scan(&aaa.Id, &aaa.Uuid, &aaa.First, &aaa.Second, &aaa.Third,
		&aaa.Email, &aaa.Password, &aaa.CreatedAt, &aaa.UpdatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		aaas = append(aaas, aaa)
	}
	rows.Close()
	return aaas, err
}

// aaaの更新
func (a *AAA) UpdateAAA() (err error) {
	cmd := `update aaa set first = ?, second = ?, third = ? where id = ?`
	_, err = Db.Exec(cmd, a.First, a.Second, a.Third, a.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err 
}