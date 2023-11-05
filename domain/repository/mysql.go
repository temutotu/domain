package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
}

type User struct {
	ID     int
	Name   string
	Pass   string
	Nation string
}

func (h *MySQL) Init() error {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/domain")
	defer db.Close()

	fmt.Println(err)
	row, err := db.Query("SELECT * FROM users")
	if err != nil {
		return err
	}

	for row.Next() {
		u := &User{}
		if err = row.Scan(&u.ID, &u.Name, &u.Nation); err != nil {
			return err
		}
		fmt.Println(u)
	}

	return err
}

func (h *MySQL) Search(name string) (string, error) {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/domain")
	defer db.Close()

	if err != nil {
		return "", err
	}

	row, err := db.Query("SELECT * FROM users WHERE name=?", name)

	u := &User{}
	for row.Next() {

		if err = row.Scan(&u.Name, &u.Pass); err != nil {
			return "", err
		}

		return u.Pass, nil
	}

	return "", nil
}

func (h *MySQL) Add(name string, pass string) error {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/domain")
	defer db.Close()

	if err != nil {
		return err
	}

	ins, err := db.Prepare("INSERT INTO users(name, pass) VALUES(?,?)")
	if err != nil {
		return err
	}

	result, err := ins.Exec(name, pass)
	fmt.Println(result)
	if err != nil {
		return err
	}

	return nil
}
