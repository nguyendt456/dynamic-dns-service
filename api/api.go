package api

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/nguyendt456/dynamic-dns-service/model"
)

func SendDDNSapi(config model.GoogleDNS, ip string) []string {
	var listOfResponse []string
	if err := net.ParseIP(ip); err == nil {
		return nil
	}
	for i := 0; i < len(config.Dns); i++ {
		if config.Dns[i].Ip == "auto" {
			config.Dns[i].Ip = string(ip)
		}
		request := fmt.Sprintf("https://%s:%s@domains.google.com/nic/update?hostname=%s&myip=%s", config.Dns[i].Username, config.Dns[i].Password, config.Dns[i].Name, ip)
		response, err := http.Get(request)
		if err != nil {
			i--
			fmt.Println(err)
			continue
		}
		content, err := io.ReadAll(response.Body)
		if err != nil {
			i--
			fmt.Println(err)
			continue
		}

		log.Printf("Request: %s \n", request)
		listOfResponse = append(listOfResponse, string(content))
		log.Printf("Response: %s \n", content)
	}
	return listOfResponse
}
