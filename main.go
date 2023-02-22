package main

import (
	"fmt"
	"sync"
	"musui/query"
	"os"
	"log"
	//"time"
)

var waitGroup sync.WaitGroup

func main()  {
	for true {
		waitGroup.Add(1)
		go GenerateIPV4()
	}
}

func GenerateIPV4() {
	defer waitGroup.Done()
	address := ""
	for true {
		n1 := "192"
		n2 := "168"
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
}