package models

import (
	// "fmt"
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string 
	Email string
	PassWord string
	CreatedAt time.Time
}

//CreateUser
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	// コマンドの実行
	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name, 
		u.Email, 
		Encrypt(u.PassWord), 
		time.Now())

	// エラーがあればエラーのログを返す
	if err != nil {
		log.Fatalln(err)
	}
	return err
}


// GetUser
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at
	from users where id = ?`

	// idを渡すためにQueryRowを使用する（ただコマンドを渡すだけではない）
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}


// UpdateUser
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`

	// 更新されたNameとEmailを渡してSQL文を実行する
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}


// DeleteUser
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}