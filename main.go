package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ncostamagna/client_request/client"
	"github.com/ncostamagna/client_request/domains/stgmachine"
)

func main() {

	myValues := []string{
		"[my Values]",
	}

	// my domain
	myDomain := &stgmachine.InitialState{
		WorkflowID:     1,
		ResourceTypeID: 1,
		StageID:        1,
	}

	myURL := "[myURL]"
	date := time.Now()
	file, err := os.Create(fmt.Sprintf("%d.log", date.UnixNano()))
	if err != nil {
		panic(err)
	}

	l := log.New(file, "", 0)
	l.Printf("URL: %s", myURL)
	for _, v := range myValues {
		myDomain.ResourceID = v
		ok, err := client.Post(myURL, "[myToken]", myDomain)
		if err != nil {
			l.SetPrefix("error: ")
			l.Printf("%s: %s", myDomain.ResourceID, err.Error())
			continue
		}

		l.SetPrefix("success: ")
		l.Printf("%s: %s", myDomain.ResourceID, string(ok))
	}
}
