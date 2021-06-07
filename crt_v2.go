package main

// Author: Eduardo Barbosa (@_anakein)
// Date: 07/06/2021
// anakein@protonmail.ch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
	fmt.Printf("%T", subdomains)
	for _, subdomain := range subdomains {
		fmt.Println("Found:", subdomain.CommonName)
	}
}

func main() {
	GetJsonFromCrt("facebook.com")
}
