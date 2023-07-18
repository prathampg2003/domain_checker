package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		checkdomain(scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Fatalf("Error :could not read from the input: %v\n", err)
	}

}

func checkdomain(domain string) {
	var haMX, hasSPF, hasDMARC bool

	var spfrecord, dmarcrecord string

	mxrecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("error in mx records")
	}

	if len(mxrecords) > 0 {
		haMX = true
	}

	textrecords, err2 := net.LookupTXT(domain)

	if err2 != nil {
		log.Printf("error in the text records")
	}

	for _, records := range textrecords {

		if strings.HasPrefix(records, "v=spf1") {
			hasSPF = true
			spfrecord = records
			break
		}
	}

	dmarrecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error%v\n", err)
	}

	for _, record := range dmarrecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcrecord = record
			break
		}
	}

	fmt.Println("domain, haMX, hasSPF, spfrecord, hasDMARC,dmarcrecord")

	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, haMX, hasSPF, spfrecord, hasDMARC, dmarcrecord)

}
