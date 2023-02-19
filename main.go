package main

import(
	"fmt"
	"flag"
	"net/http"
	"io/ioutil"
	"encoding/json"
	
)

type Crtsr struct {
	CommonName string `json:"common_name"`
	NameValue string `json:"name_value"`
}

func crt(domain string) {

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
		output = append(output, subdomains.NameValue)
		fmt.Printf("%s\n", subdomains.CommonName)
		fmt.Printf("%s\n", subdomains.NameValue)
	}
	
}


func main (){
	var p string
	flag.StringVar(&p, "p", "", "Domain name")
	flag.Parse()
	if p == ""{
		fmt.Println("Some Argument are not set")
		return
	}else {
		 crt(p)
			
	}
}

