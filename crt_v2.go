package main

// Author: Eduardo Barbosa (@_anakein)
// Date: 07/06/2021
// anakein@protonmail.ch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var banner = `
            _                
   ___ _ __| |_   __ _  ___  
 /  __| '__| __| / _  |/ _ \ 
|  (__| |  | |_ | (_| | (_) |
 \___ |_|   \__(_)__, |\___/ 
         	 |___/       
						 
 [+] by @anakein
 [+] https://github.com/an4kein
 [-] Usage: crt.go <target>
`

type Crtsr struct {
	CommonName string `json:"common_name"`
}

func GetJsonFromCrt(domain string) {

	resp, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%s&output=json", domain))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	sb := []byte(body)
	var subdomains []Crtsr
	err = json.Unmarshal(sb, &subdomains)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, subdomain := range subdomains {
		fmt.Println(subdomain.CommonName)
	}
	fmt.Println("FOUND", len(subdomains), "SUBDOMAINS")

}

func main() {
	fmt.Println(banner)
	if os.Args != nil && len(os.Args) > 1 {
		domain := os.Args[1]
		if domain != "" {
			fmt.Println("+---------------------=[Gathering Certificate Subdomains]=------------------------+")
			GetJsonFromCrt(domain)
			fmt.Println("+--------------------------------=[Done!]=----------------------------------------+")
		}
	}
}
