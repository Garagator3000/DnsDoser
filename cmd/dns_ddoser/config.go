package main

import (
	"flag"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	defaultShowConfig = false
	defaultWait       = false
	defaultCount      = 1
	defaultConfigFile = "./dns_ddos.yaml"
	defaultQuiet      = false
	defaultTimeout    = 10000
	defaultDNSServer  = "127.0.0.53:53"
	defaultHostnames  = ""
)

type Config struct {
	ShowConfig bool
	Wait       bool
	Count      int
	Quiet      bool
	Timeout    int
	DNSServer  string
	Hostnames  []string `yaml:"hostnames"`
}

func loadConfig(config *Config) error {
	help := flag.Bool("help", false, "usage")
	count := flag.Int("count", defaultCount, "count of requests to DNS server")
	wait := flag.Bool("wait", defaultWait, "press enter for continue")
	hostnames := flag.String("hosts", defaultHostnames, "hostnames for testing DNS. (--hosts ya.ru,google.com)")
	hostnamesFile := flag.String("config", defaultConfigFile, "file with test hosts")
	quiet := flag.Bool("quiet", defaultQuiet, "disable output")
	timeout := flag.Int("timeout", defaultTimeout, "timeout for dns request")
	dnsServer := flag.String("dns", defaultDNSServer, "ip:port dns server for ddos")
	showConfig := flag.Bool("show-conf", defaultShowConfig, "show configuration of Application before run")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	config.ShowConfig = *showConfig
	config.Wait = *wait
	config.Count = *count
	config.DNSServer = *dnsServer
	config.Quiet = *quiet
	config.Timeout = *timeout

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
