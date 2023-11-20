package dnsddoser

import (
	"context"
	"net"
	"time"

	"github.com/fatih/color"
)

func NewResolver(timeoutMS int, dnsServer string) *net.Resolver {
	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Duration(timeoutMS),
			}
			return d.DialContext(ctx, network, dnsServer)
		},
	}
}

func SendDnsRequest(ctx Context) ([]string, error) {
	defer ctx.Wg.Done()

	var printer *color.Color

	addrs, err := ctx.Resolver.LookupHost(context.Background(), ctx.Hostname)
	if err != nil {
		printer = color.New(color.FgRed)
		printer.Printf("cannot resolve %s: %s\n", ctx.Hostname, err)
	} else {
		ctx.Counter.Add(1)
		if ctx.Report {
			printer = color.New(color.FgGreen)
			printer.Printf("%s ----- %s\n", ctx.Hostname, addrs)
		}
	}
	return addrs, err
}
