package main

import (
	"fmt"
	"log"
	"musui/query"
	"os"
	"time"
)

// "time"

func main() {
	for true {
		go GenerateIPV4()
		time.Sleep(500 * time.Millisecond)
	}
}

func GenerateIPV4() {
	address := ""

	n1 := RandomInt(1, 255)
	n2 := RandomInt(1, 255)
	n3 := RandomInt(1, 255)
	n4 := RandomInt(1, 255)

	address = n1 + "." + n2 + "." + n3 + "." + n4 + ":19132"
	result := query.CheckServer(address)
	if result != true {
		fmt.Println("(" + address + "): NOT VALID")
	} else {
		file, err := os.OpenFile("servers.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		_, err1 := file.WriteString("(" + address + "): VALID\n")
		if err1 != nil {
			log.Fatal(err1)
		}
	}

}
