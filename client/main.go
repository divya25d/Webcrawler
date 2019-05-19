package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "github.com/tkanos/gonfig"
    "net/http"
)

type WebCrawlerConfig struct {
    CrawlUrl          string `json:"crawlUrl"`
    CrawlApi           string `json:"crawlApi"`
}

func NewWebCrawlerDefaultConfig() WebCrawlerConfig {
    return WebCrawlerConfig{
        CrawlUrl:          "http://redhat.com",
        CrawlApi:           "http://localhost:8080/crawldata/getDomainData",
    }
}

func main() {
    Cfg := WebCrawlerConfig{}
    err := gonfig.GetConf("./config.json", &Cfg)
    if err != nil {
        Cfg = NewWebCrawlerDefaultConfig()
    }
    log.Println("*****************************************************************")
    log.Println("*********************Starting Webcrawler Client******************")
    log.Println("*****************************************************************")

    client := &http.Client{}


req, err := http.NewRequest("GET", Cfg.CrawlApi, nil)
req.Header.Set("url", Cfg.CrawlUrl)
resp, err := client.Do(req)
defer resp.Body.Close()

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))

}
