package main

import (
	"flag"
	"os"
	"strings"

	dnsddoser "dns_ddoser/internal/dns-ddoser"

	"gopkg.in/yaml.v3"
)

const (
	defaultShowConfig = false
	defaultWait       = false
	defaultCount      = 0
	defaultConfigFile = "./dns_ddos.yaml"
	defaultTimeout    = 2000
	defaultDNSServer  = "127.0.0.53:53"
	defaultHostnames  = ""
	defaultLoglvl     = dnsddoser.Disable
	defaultLoglvlstr  = "info"
)

type Config struct {
	ShowConfig bool
	Wait       bool
	Count      int
	Loglevel   dnsddoser.Loglevel
	Timeout    int
	DNSServer  string
	Hostnames  []string `yaml:"hostnames"`
}

var loglvl dnsddoser.Loglevel

func loadConfig(config *Config) error {
	help := flag.Bool("help", false, "usage")
	count := flag.Int("count", defaultCount, "count of requests to DNS server")
	wait := flag.Bool("wait", defaultWait, "press enter for continue")
	hostnames := flag.String("hosts", defaultHostnames, "hostnames for testing DNS. (--hosts ya.ru,google.com)")
	hostnamesFile := flag.String("config", defaultConfigFile, "file with test hosts")
	timeout := flag.Int("timeout", defaultTimeout, "timeout for dns request")
	dnsServer := flag.String("dns", defaultDNSServer, "ip:port dns server for ddos")
	showConfig := flag.Bool("show-conf", defaultShowConfig, "show configuration of Application before run")
	loglvlstr := flag.String("loglvl", defaultLoglvlstr, "log level (disable, errors, warn, info)")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	config.ShowConfig = *showConfig
	config.Wait = *wait
	config.Count = *count
	config.DNSServer = *dnsServer
	config.Timeout = *timeout
	config.Loglevel = loglvl.Parse(*loglvlstr)

	if *hostnames != "" {
		config.Hostnames = strings.Split(*hostnames, ",")
		return nil
	}

	cf, err := os.Open(*hostnamesFile)
	if err != nil {
		return err
	}
	defer cf.Close()

	decoder := yaml.NewDecoder(cf)
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	if config.Count == defaultCount {
		config.Count = len(config.Hostnames)
	}

	return nil
}
