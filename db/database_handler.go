package database_handler

import (
	"database/sql"
	_ "github.com/lib/pq"
	"flat-scraper/flat"
	"time"
	"fmt"
	"log"
)

func FlatInsert(flats []flat.Flat) {
	db, err := sql.Open("postgres", "user=postgres password='postgres' dbname=flats")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
	  log.Fatal("Error: Could not establish a connection with the database")
	}

	fmt.Printf("Connected to DB \n")

	t := time.Now()
	t_str := fmt.Sprintf("%d-%02d-%02d",t.Year(), t.Month(), t.Day())
	for _,v := range flats {
		query := fmt.Sprintf("INSERT INTO flats(id,name,price,rooms,size,store,elevator,link,area,date_published)" +
			" VALUES (%d,'%s',%d,%d,%d,%d,%t,'%s','%s','%s');", v.Id, v.Name, v.Price, v.Rooms, v.Size, v.Store, 
			v.Elevator, v.Link, v.Area, t_str)
		db.QueryRow(query)
	}
}