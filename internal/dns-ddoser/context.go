package dnsddoser

import (
	"net"
	"sync"
	"sync/atomic"
)

type Context struct {
	Wg       *sync.WaitGroup
	Hostname string
	Resolver *net.Resolver
	Report   bool
	Counter  *atomic.Int64
}

func ContextWithResolverHostnameWG(resolver *net.Resolver, hostname string, wg *sync.WaitGroup, counter *atomic.Int64, report bool) Context {
	return Context{
		Wg:       wg,
		Hostname: hostname,
		Resolver: resolver,
		Report:   report,
		Counter:  counter,
	}

}
