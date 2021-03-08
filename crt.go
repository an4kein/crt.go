package main

//Author: Eduardo Barbosa @anakein
// 07/03/2021

// gather all certificate sub-domains from crt.sh and save them to a file
// based in https://gist.github.com/1N3/dec432d14fec84e09733f39669ebca0f


import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

var search = "https://crt.sh/?q="

var meutexto = `
           _         _     
  ___ _ __| |_   ___| |__  
 / __| '__| __| / __| '_ \ 
| (__| |  | |_ _\__ \ | | |
 \___|_|   \__(_)___/_| |_|
						 
 [+] by @anakein
 [+] https://github.com/an4kein
 [-] Usage: crt.go -domain <target>
`

func main() {
	color.Cyan(meutexto)
	color.Red("+ -- ----------------=[Gathering Certificate Subdomains]=------------------------+")
	textPtr := flag.String("domain", "", "Domain to parse. (Required!)\n Example: ./crt.go -domain example.com")
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		resp, err := http.Get(search + *textPtr)

		if err != nil {
			fmt.Println(err)
		}
		if http.StatusBadGateway == 502 {
			fmt.Println("Error: ", http.StatusBadGateway, "Bad Gateway")
		}
		//defer resp.Body.Close()
		resp.Body.Close()
		f, err := os.Create("/tmp/output.txt")
		resp.Write(f)
		//defer f.Close()
		f.Close()

	}

	cmd3 := `cat /tpm/output.txt |grep "\." |grep -v "<A style=\|.png\|href=\|&\|{\|crt.sh\|W3C/\|/2.0" |tr "<BR>" "\n"   |grep "\."  |sort -u >> /tmp/domains.txt`
	out, err := exec.Command("sh", "-c", cmd3).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
	domi, err := exec.Command("cat", "/tmp/domains.txt").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	outputDomain := string(domi[:])
	fmt.Println(outputDomain)
	color.Red("[+] Domains saved to: /tmp/domains.txt")
	color.Red("+--------------------------------=[Done!]=----------------------------------------+")

}
