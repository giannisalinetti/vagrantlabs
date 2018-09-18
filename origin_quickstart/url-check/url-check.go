// Description: url-check is a minimal tool to inspect remote urls
// Author: Giovan Battista Salinetti (gbsalinetti@extraordy.com)

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// NewProxierWithNoProxyCIDR constructs a Proxier function that respects CIDRs in NO_PROXY and delegates if
// no matching CIDRs are found.
// Copied from k8s.io/apimachinery/pkg/util/net/http.go to reflect Kubernetes behaviour.
func NewProxierWithNoProxyCIDR(delegate func(req *http.Request) (*url.URL, error)) func(req *http.Request) (*url.URL, error) {
	// we wrap the default method, so we only need to perform our check if the NO_PROXY (or no_proxy) envvar has a CIDR in it
	noProxyEnv := os.Getenv("NO_PROXY")
	if noProxyEnv == "" {
		noProxyEnv = os.Getenv("no_proxy")
	}
	noProxyRules := strings.Split(noProxyEnv, ",")

	cidrs := []*net.IPNet{}
	for _, noProxyRule := range noProxyRules {
		_, cidr, _ := net.ParseCIDR(noProxyRule)
		if cidr != nil {
			cidrs = append(cidrs, cidr)
		}
	}

	if len(cidrs) == 0 {
		return delegate
	}

	return func(req *http.Request) (*url.URL, error) {
		ip := net.ParseIP(req.URL.Hostname())
		if ip == nil {
			return delegate(req)
		}

		for _, cidr := range cidrs {
			if cidr.Contains(ip) {
				return nil, nil
			}
		}

		return delegate(req)
	}
}

func main() {

	helpFlag := flag.Bool("h", false, "Prints an help")

	flag.Parse()
	args := flag.Args()

	if !*helpFlag && len(args) == 0 {
		fmt.Println("Error: missing remote url.")
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "<REMOTE_URL>")
		os.Exit(1)
	}

	if *helpFlag {
		fmt.Println("Usage:", filepath.Base(os.Args[0]), "<REMOTE_URL>")
		os.Exit(0)
	}

	// Parse the url
	testURL := os.Args[1]
	u, err := url.Parse(testURL)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("URL Parsing section:\nScheme: %s\nHost: %s\nPath: %s\n\n", u.Scheme, u.Host, u.Path)

	// Configure insecure transport and proxies
	// Since http.ProxyFromEnvironment is used the HTTP_PROXY, HTTPS_PROXY and NO_PROXY env variables
	// are expected to work and provide proxy urls and exclusions to the client.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           NewProxierWithNoProxyCIDR(http.ProxyFromEnvironment),
	}
	client := &http.Client{Transport: tr}

	// Connect to the remote host
	resp, err := client.Get(testURL)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Print Status and Protocol along with the response body
	fmt.Printf("Response section:\nResponse: %s\nProtocol: %s\n", resp.Status, resp.Proto)
	fmt.Printf("%s", string(body))
}
