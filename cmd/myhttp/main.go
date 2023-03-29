package main

import (
	"fmt"
	"github.com/AgentNemo00/myhttp/configuration"
	"github.com/AgentNemo00/myhttp/pool"
	"log"
)

func main() {
	config, err := configuration.ParseCmdLine()
	if err != nil {
		log.Fatal(err)
	}
	p := pool.NewPool(config.Parallel)
	for _, url := range config.Urls {
		p.AddWorker(pool.WorkerByURl(url))
	}
	results := p.Do()
	for name, result := range results {
		fmt.Printf("%s %s\n", name, result)
	}

}
