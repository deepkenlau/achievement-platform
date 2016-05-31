package platformdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"os"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "e://achievement.db")
	if nil != err {
		fmt.Fprintf(os.Stderr, err.Error())
	} else {

		sqlStmt := `
		create table user(id text not null primary key, name text, status text)
		`;

		_, err = db.Exec(sqlStmt)
		if nil != err {
			fmt.Fprintf(os.Stderr, "%q: %s", err, sqlStmt)
			//return db, err
		}

		tx, err := db.Begin()
		if nil != err {
			fmt.Println(err)
		}
		stmt, err := tx.Prepare("insert into user(id, name, status) values(?, ?, ?)")
		if nil != err {
			fmt.Println(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(0, fmt.Sprintf("%s, %s, %s", "PA0001", "黄欢", "跑盘"))
		tx.Commit()

		rows, err := db.Query("select id, name, status from user")
		if nil != err {
			fmt.Println(err)
		}

		for rows.Next() {
			var id string
			var name string
			var status string

			err = rows.Scan(&id, &name, &status)
			if nil != err {
				fmt.Println(err)
			}
			fmt.Println(id, name, status)
		}
	}



	return db, err
}