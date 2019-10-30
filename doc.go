/*
Package orionxctools implements some functions to make requests 
to the Orionx api

The simplest way to use orionxtools is:
  
  package main

  import "github.com/dhinojosac/orionxclient/orionxctools"

  const (
    API_KEY = "put your API_KEY here"
    KEY     = "put your KEY here"
  )

  func main() {

    queries := []string{
      `{market(code:"CHACLP"){
        lastTrade{
          price
        }
        }
      }`,
      `{cha:marketOrderBook (marketCode:"CHACLP") {
          spread
          mid
        }}`,
    }

    for _, query := range queries {
      client.MakeRequest(KEY, API_KEY, query)
    }
  }  

Output:
  Status: 200 OK
  Body: {"data":{"market":{"lastTrade":{"price":752}}}}
  Status: 200 OK
  Body: {"data":{"cha":{"spread":20,"mid":764}}}

*/
package orionx-go-client