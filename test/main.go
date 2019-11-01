package main

// Import packets orionxctools.
import (
	"fmt"
	"log"
	"os"

	orionxctools "github.com/dhinojosac/gorionx"
)

// The KEY and API KEY were stored in environment variables.

func main() {
	//Read environmet variables
	fmt.Println("Reading Environment Variables")
	var key string
	var apiKey string

	key, ok := os.LookupEnv("ORIONX_KEY")
	if !ok {
		log.Fatal("There are not present ORIONX_KEY env var\n")
	}
	apiKey, ok = os.LookupEnv("ORIONX_API_KEY")
	if !ok {
		log.Fatal("There are not present ORIONX_API_KEY env var\n")
	}
	fmt.Printf("api: %s \napi_key: %s\n\n", key, apiKey)

	/*
		queries := []string{
			`{market(code:"CHACLP"){
				code
				name
				lastTrade{
					price
				}
			  }}`,
			`{cha:marketOrderBook (marketCode:"CHACLP") {
				spread
				mid
				}}`,
			`{marketStats(marketCode: "CHACLP", aggregation: d1, limit: 20) {
				_id
				open
				close
				high
				low
				variation
				average
				volume
				volumeSecondary
				count
				fromDate
				toDate
			  }}`,
			  		`{me{
			wallets{
				_id
				balance
				availableBalance
				unconfirmedBalance
			}
		}}`,
		}

		for _, query := range queries {
			orionxctools.MakeRequest(key, apiKey, query) // Performs API requests
		}
	*/

	/*
		//wallets
			query {
			  me {
			    wallets {
			      _id
			      balance
			      availableBalance
			      unconfirmedBalance
			       currency{
			        code
			      }

			    }
			  }
			}
	*/

	queries := []string{
		`{me{_id}}`,
		`{market(code:"CHACLP"){
			code
			name
			lastTrade{
				price
			}
		  }}`,
	}

	fmt.Println()

	for _, query := range queries {
		orionxctools.MakeRequest(key, apiKey, query) // Performs API requests
	}

	/*
		//market order
		mutations := []string{
			`mutation {placeMarketOrder(marketCode:"CHACLP", amount:520, sell: false){
				_id
				__typename
			  }}`,
		}
	*/

	//limit order
	mutations := []string{
		`mutation {placeLimitOrder(marketCode:"CHACLP", amount:1400000000, limitPrice:40, sell:true){
				_id
				__typename
			  }}`,
	}

	/*
		//Cancel order by id
		mutations := []string{
			`mutation {cancelOrder(orderId:"mf248tMuoFztzt7Jg"){
				_id
				__typename
				status
			  }}`,
		}
	*/

	for _, mutation := range mutations {
		orionxctools.MakeRequest(key, apiKey, mutation) // Performs API requests
	}

}
