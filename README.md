Searches for abebooks via the search api, this allows you to search for books using a few filters, other filter options 
are available but need to be extended in this repo

The official abebooks documentation can be found in this folder 

### Example usage

Search for a book by ISBN

```
package main

import (
	"github.com/mkadiri/abebooks-search-sdk/model"
	"github.com/mkadiri/abebooks-search-sdk/service"
	"fmt"
)

func main() {
	client := service.AbebooksClient{
		ClientKey: "API_KEY",
		Currency: "GBP",
		DestinationCountry: "UK",
		VendorLocation: "UK",
	}

	filter := model.AbebooksSearchFilter{
		BookCondition: model.ConditionUsed,
		OutputSize: model.OutputSizeLong,
	}

	results, err := client.SearchByIsbn("9781408845653", filter)

	if err != nil {
		fmt.Print("fail: " + err.Error())
		return
	}

	fmt.Printf("%+v\n", results)
}

```