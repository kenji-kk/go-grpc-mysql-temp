package db

import (
	"log"
)

type Country struct {
	Id string
	Name string
}

func GetAllCountry() (allCountry []Country, err error) {
	cmd := `SELECT * FROM country`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		country := Country{}
		err = rows.Scan(&country.Id, &country.Name)
		if err != nil {
			log.Fatalln(err)
		}
		allCountry = append(allCountry, country)
	}
	rows.Close()

	return allCountry, nil
}
