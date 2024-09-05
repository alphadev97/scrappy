package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	var url string

	fmt.Println("Please type website url: ")

	_, err := fmt.Scanln(&url)
	if err != nil {
		fmt.Println("Error reading url: ", err)
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching url: ", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error in reading body: ", err)
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Println("Error in parsing HTML: ", err)
	}

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			if n.FirstChild != nil {
				fmt.Printf("%s: %s \n", n.Data, n.FirstChild.Data)
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
}
