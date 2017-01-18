package database_handler

import (
	"database/sql"
	_ "github.com/lib/pq"
	"flat-scraper/flat"
	"time"
	"fmt"
	"log"
)

func FlatInsert(flats []flat.Flat) []flat.Flat {
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
	
	var newFlats []flat.Flat
	
	checkQuery, err := db.Prepare("select exists(select 1 from flats where id=$1) AS \"exists\";")
	if err != nil {
	  log.Fatal(err)
	}

	insertQuery, err := db.Prepare("INSERT INTO flats(id,name,price,rooms,size,store,elevator,link,area,date_published)" +
			" VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);")
	if err != nil {
	  log.Fatal(err)
	}

	var checkRes bool
	for _,v := range flats {
		
		err = checkQuery.QueryRow(v.Id).Scan(&checkRes)
		
		if (!checkRes) {
		  insertQuery.Exec(v.Id, v.Name, v.Price, v.Rooms, v.Size, v.Store, v.Elevator, v.Link, v.Area, t_str)
		  
		  newFlats = append(newFlats,v)
		}
		
	}
	
	return newFlats
}