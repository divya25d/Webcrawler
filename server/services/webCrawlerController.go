package services

import (
	"../types"
	"crypto/tls"
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
	"net/url"

	"gopkg.in/yaml.v2"
	"strings"
)

var visited = make(map[string]bool)

//To get domain details
func GetDomainDetails(uri string) ([]byte, error) {
	fmt.Println("fetching", uri)
	visited[uri] = true
	linksData := types.DomainData{}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{Transport: transport}
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	links := collectlinks.All(resp.Body)

	for _, link := range links {
		absolute := fixUrl(link, uri)
		if strings.HasPrefix(absolute, uri) {
			if !visited[absolute] {
				linksData.Links = append(linksData.Links, absolute)
			}
		}

	}
	data, err := yaml.Marshal(linksData)
	if err != nil {
		return nil, err
	}
	return data, err

}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}
