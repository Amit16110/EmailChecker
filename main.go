package main

import (
	"bufio" // the buffer package to be able to parse whatever i put in terminal.
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")
	fmt.Println("Input here")

	for scanner.Scan(){
		checkDomain(scanner.Text()) //putting the domain name one by one
	}
	if err:= scanner.Err(); err != nil {
		log.Printf("Error: could not read from input: %v\n", err) 
	}
}
//it take the domain name as an argu
func checkDomain(domain string){

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _,record := range txtRecords{
		if strings.HasPrefix(record,"v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err	:= net.LookupTXT("_dmarc." + domain)
		if err!=nil {
			log.Printf("Error: %v\n", err)
		}
		for _,record := range dmarcRecords{
			if strings.HasPrefix(record,"v=DMARC1") {
				hasDMARC = true
				dmarcRecord = record
				break
			}
		}

		fmt.Printf("Domain-Name:=> %v,\nHasMX:=> %v,\nHasSPF:=>%v,\nSPRRecord:=>%v,\nHasDMARC:=>%v,\nDmarcRecord:=>%v\n",  domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}