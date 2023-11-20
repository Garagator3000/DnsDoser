all: build

build:
	go build ./cmd/dns_ddoser/

clean:
	rm dns_ddoser