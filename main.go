package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	file, err := os.Create("scraped_content.txt")
	if err != nil {
		log.Fatal("Could not create file", err)
	}

	defer file.Close()

	c := colly.NewCollector(
		colly.AllowedDomains("rojrztech.com"),
	)

	vistied := make(map[string]bool)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		_, err := file.WriteString(fmt.Sprintf("Visiting: %s\n", r.URL.String()))
		if err != nil {
			log.Fatal("Could not write the file", err)
		}
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		fmt.Println("Paargraph: ", e.Text)
    _, err := file.WriteString(fmt.Sprintf("Paargraph: %s\n", e.Text))
    if err != nil {
      log.Fatal("Could not write paragraph", err)
    }
	})

  c.OnHTML("a[href]", func(e *colly.HTMLElement){
    link := e.Attr("href")
    if !vistied[link] {
      vistied[link] = true

      err := c.Visit(e.Request.AbsoluteURL(link))
      if err != nil {
        fmt.Println("Error visiting link", err)
      }
    }
  })

  err = c.Visit("http://rojrztech.com")
  if err != nil {
    fmt.Println("Error:", err)
  }

}
