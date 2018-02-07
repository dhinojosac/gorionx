
# orionxctools 

This package was created to facilitate queries to the oriox API.


#### Example

The simplest way to use orionxtools is:

```go
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
```