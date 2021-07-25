package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	username, queryErr := query(db, 123)

	if queryErr != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(queryErr), errors.Cause(queryErr))
		fmt.Printf("stack trace: \n%+v\n", queryErr)
		// 返回业务错误信息
		fmt.Printf("No user with that ID. \n")
	}
	// 返回正确数据
	fmt.Printf("Username is %s\n", username)
}

func query(db *sql.DB, id int) (string, error) {
	var username string
	scanErr := db.QueryRow("SELECT username FROM users WHERE id=?", id).Scan(&username)
	// 应该把 ErrNoRows 错误 wrap 后返回，因为 dao 层不确定业务逻辑
	switch {
	case scanErr == sql.ErrNoRows:
		// log.Printf("No user with that ID.")
		return "", errors.Wrap(scanErr, "No data with that ID.")

	default:
		return username, nil
	}
}
