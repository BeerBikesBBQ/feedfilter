package main

import (
  "flag"
  "fmt"
  "github.com/mmcdole/gofeed"
)

func main() {

  var url string

  flag.StringVar(&url,"url","","The feed URL")
  flag.Parse()

  fmt.Println("URL to parse:", url)

  fp := gofeed.NewParser()
  feed, _ := fp.ParseURL(url)

  fmt.Println("Feed title:", feed.Title)
}
