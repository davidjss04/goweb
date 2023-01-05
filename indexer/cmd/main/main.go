package main

import (
	"fmt"

	utils "github.com/davidjss04/indexer/pkg/utils"

	// "fmt"
	// "net/http"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"

)

func main() {
	// start := time.Now()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Go server running on port 6060")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go leakyFunction(wg, func() {
		http.ListenAndServe("localhost:6060", nil)
		wg.Done()
		log.Println("Server stopped")
	})
	wg.Wait()

	// wg := &sync.WaitGroup{}
	// wg.Add(1)

	// // go utils.LoadMails("/home/davidjss/workspace/dataChallenge/tree", "emails", wg)
	// go utils.LoadMails("/home/davidjss/workspace/dataChallenge/enron_mail_20110402/maildir", "emailsV1", wg)
	// wg.Wait()
	// log.Printf("Execution time %s\n", time.Since(start))
}

func leakyFunction(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go f()
	// go utils.LoadMails("/home/davidjss/workspace/dataChallenge/tree", "emails")
	go utils.LoadMails("/home/davidjss/workspace/dataChallenge/enron_mail_20110402/maildir", "emails", wg)

	wg.Wait()
}
