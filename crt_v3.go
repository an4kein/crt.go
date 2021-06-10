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

func GetJsonFromCrt(domain string) ([]string, error) {

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
	output := make([]string, 0)
	for _, subdomains := range subdomains {
		output = append(output, subdomains.CommonName)
	}
	return output, nil

}
func removeDuplicateValues(strSlice []string) []string {
	// https://www.geeksforgeeks.org/how-to-remove-duplicate-values-from-slice-in-golang/
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	fmt.Println(banner)
	if os.Args != nil && len(os.Args) > 1 {
		domain := os.Args[1]
		if domain != "" {
			fmt.Println("+---------------------=[Gathering Certificate Subdomains]=------------------------+")
			subdom, err := GetJsonFromCrt(domain)
			if err != nil {
				fmt.Println("error: ", err)
			}
			removeDuplicateValuesSlice := removeDuplicateValues(subdom)

			// Printing the filtered slice
			// without duplicates
			for _, uniquesubdomain := range removeDuplicateValuesSlice {
				fmt.Println(uniquesubdomain)
			}
			fmt.Println("+--------------------------------=[Done!]=----------------------------------------+")
		}
	}
}
