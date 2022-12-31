package main

import (
	utils "github.com/davidjss04/indexer/pkg/utils"

	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	// http.HandleFunc("/debug/pprof/", pprof.Index)
	// http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// http.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// utils.LoadMails("/home/davidjss/workspace/zincsearch/indexer/tree", "export.json", "mailsv1")

    // we need a webserver to get the pprof webserver
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    fmt.Println("Go server running on port 6060")
    var wg sync.WaitGroup
    wg.Add(1)
    go leakyFunction(wg)
    wg.Wait()
}

func leakyFunction(wg sync.WaitGroup) {
	defer wg.Done()


	for i := 0; i < 10; i++ {
		utils.LoadMails("/home/davidjss/workspace/zincsearch/enron_mail_20110402/maildir", "export.json", "newMails")  
	}


}


