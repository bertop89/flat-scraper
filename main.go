package main

import (
  "log"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  "flat-scraper/flat"
  "flat-scraper/utils"
  "flat-scraper/db"
  "fmt"
  "sync"
  "bytes"
)

func FlatScraper(c utils.Configuration) []flat.Flat {

  flatResponses := make(chan []flat.Flat)

  var wg sync.WaitGroup

  wg.Add(len(c.Areas))

  for _, area := range c.Areas {
    go func(area string) {
      fmt.Println(area+" starts")
      var resultList []flat.Flat
      doc, err := goquery.NewDocument("https://www.idealista.com/alquiler-viviendas/madrid/"+area+"/con-pisos,estado_buen-estado,amueblado_amueblados/") 
      if err != nil {
        log.Fatal(err)
      }

      doc.Find(".items-container article").Each(func(i int, s *goquery.Selection) {
        name := s.Find("a").Text()
        price, err := strconv.Atoi(strings.Replace(strings.Split(s.Find(".item-price").Text(),"€")[0],".","",-1))
        var rooms, size, store, id int
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
        if strings.Contains(link, "/inmueble/") {
            id, err = strconv.Atoi(strings.Split(link, "/")[2])
            var current = flat.Flat{id, name, price, rooms, size, store, elevator, link, area}
            resultList = append(resultList,current)
        }
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
  }

  return finalList

  
}

func Filter(vs []flat.Flat, c utils.Configuration) []flat.Flat {
    var vsf []flat.Flat
    for _, v := range vs {
        if v.Price<= c.Price && 
           v.Size >= c.Size &&
           v.Elevator && v.Store > 1 {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func main() {
  config := utils.LoadConfig()

  resultList := FlatScraper(config)
  
  newFlats := database_handler.FlatInsert(resultList)
  
  filteredList := Filter(newFlats, config)

  if (len(filteredList) > 0) {
  
    filteredLength := strconv.Itoa(len(filteredList))
    var buffer bytes.Buffer
    for _,v := range filteredList {
      buffer.WriteString(v.ToString()+"\n")
    }
    filteredListText := buffer.String()
    utils.EmailSend(config, filteredLength, filteredListText)
    
  }
}
