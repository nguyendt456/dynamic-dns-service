package test

import (
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/nguyendt456/dynamic-dns-service/api"
	"github.com/nguyendt456/dynamic-dns-service/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

type APITestSuite struct {
	suite.Suite
	GoogleDNS model.GoogleDNS
	Config    model.Config
}

func (s *APITestSuite) SetupSuite() {
	log.Println("Setting up...")
}

func (s *APITestSuite) Test1_ReadConfigFile() {
	content, err := ioutil.ReadFile("config_test.yaml")
	if err != nil {
		s.T().Fatal(err)
	}

	var dns model.GoogleDNS
	var provider model.Config
	err = yaml.Unmarshal(content, &dns)
	err = yaml.Unmarshal(content, &provider)

	assert.Equal(s.T(), "google", provider.Provider)
	assert.Equal(s.T(), "ddnstest.binhnguyen.dev", dns.Dns[0].Name)

	s.Config = provider
	s.GoogleDNS = dns
}

func (s *APITestSuite) Test2_SendAPI() {
	log.Println("Sending DNS api...")
	listOfResponse := api.SendDDNSapi(s.GoogleDNS, "1.2.3.4")
	if listOfResponse == nil {
		log.Fatal("Failed reason: no response found")
	}
	for i, v := range listOfResponse {
		if i == 0 {
			assert.Equal(s.T(), "good 1.2.3.4", v)
		}
	}
	time.Sleep(time.Second * 2)
}

func (s *APITestSuite) TearDownSuite() {
	log.Println("Sending DNS api to reset ...")
	listOfResponse := api.SendDDNSapi(s.GoogleDNS, "0.0.0.0")
	if listOfResponse == nil {
		log.Fatal("Failed reason: no response found")
	}
	for i, v := range listOfResponse {
		if i == 0 {
			assert.Equal(s.T(), "good 0.0.0.0", v)
		}
	}
}

func TestMain(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
