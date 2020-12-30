package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"../config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	// tableNamePeople = "people"
	tableNameTodo = "todos"
	tableNameA = "aaa"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	//Userテーブルがなければ作成する
	cmdU := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincrement,
		uuid string not null unique,
		name string,
		email string,
		password string,
		created_at datetime)`, tableNameUser)
	//コマンドを実行する
	Db.Exec(cmdU)

	// cmdP := fmt.Sprintf(`create table if not exists %s(
	// 	id integer primary key autoincrement,
	// 	uuid string not null unique,
	// 	firstname string, 
	// 	lastname string, 
	// 	email string unique,
	// 	password string,
	// 	created_at datetime,
	// 	updated_at datetime 
	// )`, tableNamePeople)
	// Db.Exec(cmdP)

	cmdT := fmt.Sprintf(`create table if not exists %s (
		id integer primary key autoincrement,
		content text user_id integer,
		user_id integer,
		created_at datetime)`, tableNameTodo)
	Db.Exec(cmdT)

	cmdA := fmt.Sprintf(`create table if not exists %s(
		id integer primary key autoincreament,
		uuid string not null unique,
		first string,
		second string,
		third string,
		email string,
		password string,
		created_at datetime,
		updated_at datetime)`, tableNameA)
	Db.Exec(cmdA)
}

// UUIDのパッケージを使用してUUIDを作成する
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// パスワードをハッシュ値にする
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}

