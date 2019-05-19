package main

import (
	"./httplib"
	"github.com/tkanos/gonfig"
	"log"
)

type WebCrawlerConfig struct {
	HTTPPort          string `json:"httpPort"`
}

func NewWebCrawlerDefaultConfig() WebCrawlerConfig {
	return WebCrawlerConfig{
		HTTPPort:          "80",
	}
}

func main() {
	Cfg := WebCrawlerConfig{}
	err := gonfig.GetConf("./config.json", &Cfg)
	if err != nil {
		Cfg = NewWebCrawlerDefaultConfig()
	}
	log.Println("*****************************************************************")
	log.Println("*********************Starting Webcrawler Service******************")
	log.Println("*****************************************************************")
	httplib.Run(Cfg.HTTPPort)
	wait := make(chan struct{})
	//Wait for termination signal
	<-wait
}
