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
	Db, err := sql.Open("mysql", "go_grpc:password@tcp(mysql:3306)/go_database?charset=utf8parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("db接続完了")

	cmdC := `CREATE TABLR IF NOT EXISTS country {
		id int(11) NOT NULL AUTO_INCREMENT,
		name varchar(255))`

		_, err = Db.Exec(cmdC)
		count := 0
		if err != nil {
			for {
				if err == nil {
					fmt.Println("CREATE成功")
					break
				}
				fmt.Print(".")
				time.Sleep(time.Second)
				count++
				if count > 100 {
					fmt.Println("")
					panic(err)
				}
				_, err = Db.Exec(cmdC)
			}
		}

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
