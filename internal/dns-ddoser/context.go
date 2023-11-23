package dnsddoser

import (
	"net"
	"sync"
	"sync/atomic"
)

type Loglevel int

const (
	Disable = iota
	Errors
	Warn
	Info
)

func (loglvl *Loglevel) Parse(loglvlstr string) Loglevel {
	switch {
	case loglvlstr == "disable":
		return Disable
	case loglvlstr == "errors":
		return Errors
	case loglvlstr == "warn":
		return Warn
	case loglvlstr == "info":
		return Info
	default:
		return Disable
	}
}

type Context struct {
	Wg       *sync.WaitGroup
	Hostname string
	Resolver *net.Resolver
	Loglevel Loglevel
	Counter  *atomic.Int64
}

func ContextWithResolverHostnameWG(resolver *net.Resolver, hostname string, wg *sync.WaitGroup, counter *atomic.Int64, loglvl Loglevel) Context {
	return Context{
		Wg:       wg,
		Hostname: hostname,
		Resolver: resolver,
		Loglevel: loglvl,
		Counter:  counter,
	}

}
