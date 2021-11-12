package main

//go:generate go generate ./ent

import (
	"context"
	"ent-sample/ent"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	db, err := ent.Open("postgres", "host=localhost port= user= dbname= password= sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// run the auto migration tool.
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	user := db.User.Create()
	user.SetPasswordHash("hoge")
	user.SetUserType(1)
	_, err = user.Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	staff := db.Staff.Create()
	staff.SetEmail("hoge@example.com")
	staff.SetFirstName("fname")
	staff.SetLastName("lname")
	staff.SetRole(1)
	staff.SetGender(1)
	staff.SetUserId(uuid.New())
	_, err = staff.Save(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
