package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error 

func DbConnect() {
	Db, err = sql.Open("mysql", "go_grpc:password@tcp(mysql:3306)/go_database?charset=utf8&parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("db接続完了")

	cmdC := `CREATE TABLE IF NOT EXISTS country (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255))`

		_, err = Db.Exec(cmdC)
		count := 0
		if err != nil {
			for {
				if err == nil {
					fmt.Println("")
					break
				}
				fmt.Print(".")
				time.Sleep(time.Second)
				count++
				if count > 180 {
					fmt.Println("")
					panic(err)
				}
				_, err = Db.Exec(cmdC)
			}
		}

		fmt.Println("テーブル作成成功")

		cmd := `INSERT INTO country(name) VALUES(?)`
		_, err = Db.Exec(cmd, "Japan")
		if err != nil {
			log.Fatalln((err))
		}

		cmd = `INSERT INTO country(name) VALUES(?)`
		_, err = Db.Exec(cmd, "China")
		if err != nil {
			log.Fatalln((err))
		}

		cmd = `INSERT INTO country(name) VALUES(?)`
		_, err = Db.Exec(cmd, "America")
		if err != nil {
			log.Fatalln((err))
		}
}
