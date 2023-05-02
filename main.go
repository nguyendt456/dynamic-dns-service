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

	"gopkg.in/yaml.v3"
)

type Config struct {
	Provider string `yaml:"provider"`
}

type GoogleDNS struct {
	Config
	Dns []DNS `yaml:"dns"`
}

type DNS struct {
	Name     string `yaml:"name"`
	Ip       string `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func sendDDNSapi(config GoogleDNS, ip string) {
	if err := net.ParseIP(ip); err == nil {
		return
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
		log.Printf("Response: %s \n", content)
	}
}
func checkIPonRemoteDomain(config GoogleDNS, newIP string) {
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
				sendDDNSapi(config, newIP)
				ioutil.WriteFile("ip", []byte(newIP), 0644)
				time.Sleep(time.Second)
			} else {
				return
			}
		}

	}
}

func checkIP(config GoogleDNS) {
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
			go sendDDNSapi(config, string(newIP))
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

	var result GoogleDNS
	var provider Config
	err = yaml.Unmarshal(content, &result)
	err = yaml.Unmarshal(content, &provider)
	if err != nil {
		log.Fatal(err)
	}

	if provider.Provider == "google" {
		checkIP(result)
	}
}
