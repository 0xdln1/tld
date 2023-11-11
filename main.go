package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func extractDomainAndTLD(subdomain string) (string, string, error) {
	// Use the publicsuffix package to extract TLD
	domain, err := publicsuffix.EffectiveTLDPlusOne(subdomain)
	if err != nil {
		return "", "", err
	}
	return domain, subdomain, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read subdomains from stdin
	for scanner.Scan() {
		subdomain := strings.TrimSpace(scanner.Text())
		if subdomain == "" {
			continue
		}

		domain, _, err := extractDomainAndTLD(subdomain)
		if err != nil {
			log.Printf("Error extracting domain and TLD for subdomain %s: %v", subdomain, err)
		} else {
			fmt.Println(domain)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading from stdin:", err)
	}
}
