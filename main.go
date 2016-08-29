package main

import (
  "log"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "flat-scraper/flat"
  "flat-scraper/utils"
  "fmt"
  "sync"
)

func FlatScraper(c utils.Configuration) []flat.Flat {

  

  flatResponses := make(chan []flat.Flat)

  var wg sync.WaitGroup

  wg.Add(len(c.Areas))

  for _, area := range c.Areas {
    go func(area string) {
      fmt.Println(area+" starts")
      //defer wg.Done()
      var resultList []flat.Flat
      doc, err := goquery.NewDocument("https://www.idealista.com/alquiler-viviendas/madrid/"+area+"/con-pisos,estado_buen-estado,amueblado_amueblados/") 
      if err != nil {
        log.Fatal(err)
      }

      // Find the review items
      doc.Find(".items-container article").Each(func(i int, s *goquery.Selection) {
        // For each item found, get the band and title
        name := s.Find("a").Text()
        price, err := strconv.Atoi(strings.Replace(strings.Split(s.Find(".item-price").Text(),"€")[0],".","",-1))
        var rooms, size, store int
        var elevator bool = false
        s.Find(".item-detail").Each(func(i int, sel *goquery.Selection) {
          item := sel.Text()
          if strings.Contains(item, "hab") {
            rooms, err = strconv.Atoi(strings.Split(item, " hab")[0])
          } else if strings.Contains(item, "m") {
            size, err = strconv.Atoi(strings.Split(item, " m")[0])
          } else if strings.Contains(item, "ª") {
            store, err = strconv.Atoi(strings.Split(item, "ª")[0])
            if strings.Contains(item, "con") {
              elevator = true
            }
          }
        })
        link := s.Find("a").AttrOr("href", "")
        var current = flat.Flat{name, price, rooms, size, store, elevator, link}
        resultList = append(resultList,current)
      })
      fmt.Printf("%s finishes with size %d \n",area,len(resultList))
      flatResponses <- resultList
      wg.Done()
    }(area)
  }

  go func() {
    wg.Wait()
    close(flatResponses)
  }()

  var finalList []flat.Flat

  for response := range flatResponses {
      finalList = append(finalList, response...)
      fmt.Println("receiving data... size %d", len(finalList))
  }

  return finalList

  
}

func Filter(vs []flat.Flat) string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if v.Price<= 650 && v.Size >= 50 {
            vsf = append(vsf, v.ToString())
        }
    }
    return strings.Join(vsf,"")
}

func main() {
  config := utils.LoadConfig()
  resultList := FlatScraper(config)
  filteredListText := Filter(resultList)
  utils.EmailSend(config, filteredListText)
}
