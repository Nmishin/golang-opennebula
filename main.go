package main

import (
	"log"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
)

func main() {
	con := map[string]string{
		"user":     "user_name",
		"password": "user_password",
		"endpoint": "opennebula xml-rpc endpoint",
	}

	client := goca.NewDefaultClient(
		goca.NewConfig(con["user"], con["password"], con["endpoint"]),
	)

	controller := goca.NewController(client)

	users_controller := controller.Users()

	user_id, err := users_controller.ByName(con["user"])
	if err != nil {
		log.Fatalf("Failed to get user ID from OpenNebula: %v\n", err)
	}

	log.Printf("UserID: %v\n", user_id)
}
