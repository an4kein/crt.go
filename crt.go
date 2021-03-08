package main

//Author: Eduardo Barbosa @anakein
// 07/03/2021

// gather all certificate sub-domains from crt.sh and save them to a file
// based in https://gist.github.com/1N3/dec432d14fec84e09733f39669ebca0f

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/fatih/color"
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

var search = "https://crt.sh/?q="

func crtgo() {
	if os.Args != nil && len(os.Args) > 1 {
		domain := os.Args[1]
		if domain != "" {
			color.Red("+---------------------=[Gathering Certificate Subdomains]=------------------------+")
			resp, err := http.Get(search + domain)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			f, err := os.Create("/tmp/output.txt")
			defer f.Close()
			resp.Write(f)
			cmd3 := `cat /tmp/output.txt |grep "\." |grep -v "<A style=\|.png\|href=\|&\|{\|crt.sh\|W3C/\|/2.0" |tr "<BR>" "\n"   |grep "\."  |sort -u >> /tmp/domains.txt `
			exec.Command("sh", "-c", cmd3).Output()
			cmd4 := `cat /tmp/domains.txt \n; wc -l /tmp/domains.txt`
			out2, err := exec.Command("sh", "-c", cmd4).Output()
			if err != nil {
				fmt.Printf("%s", err)
			}
			output2 := string(out2[:])
			fmt.Println(output2)
			color.Green("[+] Domains saved to: /tmp/domains.txt")
			color.Red("+--------------------------------=[Done!]=----------------------------------------+")
		} else {
			domain = ""
			fmt.Println("No command given")
		}
	}

}
func main() {
	color.Cyan(banner)
	crtgo()
}
