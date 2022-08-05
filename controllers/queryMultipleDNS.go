package controllers

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"golang.org/x/net/idna"
)

func QueryMultipleDNS(domainName string) string {
	nameservers := map[string]string{
		"Google": "8.8.8.8:53",
		//"Google IPv6":     "[2001:4860:4860::8888]:53",
		"Cloudflare": "1.1.1.1:53",
		//"Cloudflare IPv6": "[2606:4700:4700::1111]:53",
		"OpenDNS": "208.67.222.222:53",
		//"OpenDNS IPv6":    "[2620:119:35::35]:53",
		"Freifunk MUC":         "5.1.66.255:53",   //IPv6: 2001:678:e68:f000::"
		"Censurfridns Denmark": "89.233.43.71:53", //IPv6: 2001:67c:28a4::
	}

	domainNamePunycode, err := idna.Display.ToASCII(domainName)
	if err != nil {
		return "No valid domain name"
	}

	output := ""
	for nsName, nsIP := range nameservers {
		output += resolveAndPrint(nsName, nsIP, domainNamePunycode) + "\n"
	}

	return output
}

func resolveAndPrint(nsName, nsIP, domainName string) string {
	output := fmt.Sprintf("Resolving %s with %20s: ", domainName, nsName)
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, nsIP)
		},
	}
	resolvedIPs, err := r.LookupHost(context.Background(), domainName)
	if err != nil {
		output += strings.Split(err.Error(), ": ")[1]
		return output
	}
	for _, ip := range resolvedIPs {
		output += ip + " "
	}
	return output
}
