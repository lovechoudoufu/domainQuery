package internal

import (
	"context"
	"encoding/csv"
	"net"
	"strings"
	"sync"
	"time"
)

// 解析cname
func QueryCname(domain string, dnsServer string) string {
	//dnsServer := "114.114.114.114:53"
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second * 3,
			}
			return d.DialContext(ctx, network, dnsServer)
		},
	}

	cname, err := resolver.LookupCNAME(context.Background(), domain)
	if err != nil {
		return ""
	}

	cname = strings.TrimSuffix(cname, ".")
	if cname == domain {
		return ""
	}

	return cname
}

// 解析A
func QueryA(domain string, dnsServer string) []string {
	//dnsServer := "114.114.114.114:53"
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second * 3,
			}
			return d.DialContext(ctx, network, dnsServer)
		},
	}

	ips, err := resolver.LookupIP(context.Background(), "ip", domain)
	if err != nil {
		return nil
	}

	var ipStrings []string
	for _, ip := range ips {
		ipStrings = append(ipStrings, ip.String())
	}

	return ipStrings
}

// 获取域名解析结果
func QueryDomianRecord(mywrite *csv.Writer, domain string, dnsServer string, mu *sync.Mutex) {
	// cname的域名先解析出cname是什么，再解析ip是什么
	cname := QueryCname(domain, dnsServer)
	if cname != "" {
		ipStrings := QueryA(domain, dnsServer)
		for _, ip := range ipStrings {
			Search(domain, cname, mywrite, ip, mu)
		}
		//fmt.Println("[+] " + domain)
		//WriteCSV(mywrite, domain, cname, "", "", "", mu)
	} else {
		// A的域名直接获取IP地址
		ipStrings := QueryA(domain, dnsServer)
		if len(ipStrings) == 0 {
			WriteCSV(mywrite, domain, "", "", "", "", mu)
		}
		for _, ip := range ipStrings {
			Search(domain, "", mywrite, ip, mu)
			//fmt.Println("[+] " + domain)
		}
	}
}

// isIPAddress 函数用于判断字符串是否为IPv4或IPv6地址
func IsIPAddress(s string) bool {
	return net.ParseIP(s) != nil
}
