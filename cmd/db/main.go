package main

import (
	"fmt"
	"go-replica-db/pkg/db"
)

func main() {
	for i := range 3 {
		err := db.Insert("leader.txt", "testKey", fmt.Sprintf("testValue%d", i))
		if err != nil {
			print(err.Error())
			return
		}
	}

	val, err := db.Read("data/logs/leader.txt", "testKey")
	if err != nil {
		print(err.Error())
		return
	}
	print(val)
}
