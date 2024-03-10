package main

import (
	"fmt"
)

var domain string = "site.com"
var locale string = ""

func DomainForLocale(domain, locale string) string {
	subdomain := ""
	if locale == "" {
		subdomain = "en"
	} else {
		subdomain = locale
	}

	return fmt.Sprintf("%s.%s", subdomain, domain)
}

func main() {
	fmt.Println(DomainForLocale(domain, locale))
}
