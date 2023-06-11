# go-mdc
Missouri Department of Conservation data parsing library for Go

## Waterfowl Harvest Data

The library is able to request and parse waterfowl harvest data from the [MDC website](https://extra.mdc.mo.gov/widgets/wtrfwl_harvest/).

## Install

Run the following command to get the library for your project

```bash
go get github.com/lum8rjack/go-mdc
```

## Example

Using the library to get 2022 harvest data for Ten Mile Pond.

```go
package main

import (
	"fmt"
	"log"

	"github.com/lum8rjack/go-mdc"
)

func main() {
	// Select the area to get data for
	tenmile, err := mdc.GetAreaByName("ten mile")
	if err != nil {
		log.Fatalln(err)
	}

	// Get data from 2022
	data, err := mdc.GetData(2022, tenmile)
	if err != nil {
		log.Fatalln(err)
	}

	// Get the first day
	firstday := data.WfHrvUpdatesRecs[0]
	fmt.Printf("Date: %s\n", firstday.ReportDate)
	fmt.Printf("Ducks harvested: %d\n", firstday.DuckHarv)
}
```

Output:

```bash
Date: 2022-11-19T00:00:00
Ducks harvested: 40
```

