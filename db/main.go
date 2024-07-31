package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to database successfully")

	userRepo := NewUserRepository(db)

	users, err := userRepo.GetAll()
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("GET ALL:", users)

	newUser := User{
		Username: "doe",
		Age: 38,
	}
	err = userRepo.Insert(&newUser)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("INSERT:", newUser)

	user, err := userRepo.GetByID(newUser.ID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("GET BY ID:", user)

	user.Username = "ron"
	user.Age = 27
	err = userRepo.Update(user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("UPDATE:", user)

	var id int64 = 10
	err = userRepo.DeleteByID(id)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println("DELETE:", id)
}

func openDB() (*sql.DB, error) {
	driver := "pgx"
	dsn := "postgres://postgres:pass1234@localhost:5432/testdb"

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
