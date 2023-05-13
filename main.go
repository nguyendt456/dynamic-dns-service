package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/nguyendt456/dynamic-dns-service/api"
	"github.com/nguyendt456/dynamic-dns-service/model"
	"gopkg.in/yaml.v3"
)

func checkIPonRemoteDomain(config model.GoogleDNS, newIP string) {
	for i := 0; i < len(config.Dns); i++ {
		ip, err := net.LookupIP(config.Dns[i].Name)
		if err != nil {
			i--
			continue
		}
		for _, ips := range ip {
			if !(ips.String() == newIP) {
				fmt.Println(ips, newIP)
				fmt.Println("Retrying when not matching with remote")
				api.SendDDNSapi(config, newIP)
				ioutil.WriteFile("ip", []byte(newIP), 0644)
				time.Sleep(time.Second)
			} else {
				return
			}
		}

	}
}

func checkIP(config model.GoogleDNS) {
	prevIP, err := ioutil.ReadFile("ip_cache")
	if err != nil {
		os.Create("ip_cache")
		prevIP, _ = ioutil.ReadFile("ip_cache")
	}
	for {
		log.Println("Requesting...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", "http://ipinfo.io/ip", nil)
		if err != nil {
			fmt.Println(err)
			continue
		}
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out, retrying...")
			continue
		}

		newIP, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if err := net.ParseIP(string(newIP)); string(prevIP) != string(newIP) && err != nil {
			go api.SendDDNSapi(config, string(newIP))
			ioutil.WriteFile("ip_cache", newIP, 0644)
			prevIP = newIP
		} else {
			log.Println("nochange")
		}
		fmt.Println("newIP: ", string(newIP))
		checkIPonRemoteDomain(config, string(newIP))

		time.Sleep(time.Second * 2)
	}
}

func main() {
	content, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var result model.GoogleDNS
	var provider model.Config
	err = yaml.Unmarshal(content, &result)
	err = yaml.Unmarshal(content, &provider)
	if err != nil {
		log.Fatal(err)
	}

	if provider.Provider == "google" {
		checkIP(result)
	}
}
