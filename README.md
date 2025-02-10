# ClearDate

Everyone loves reference-based time formatting so much. That I had to create this package for fun. Please don't take it seriously, in the best case you don't even use it. But if you do, I understand you...

## Installation

```bash
go get github.com/MisterJD/go-cleardate
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/MisterJD/go-cleardate"
)

func main() {
	now := time.Now()
	customLayout := "EEEE, MMMM d, yyyy hh:mm:ss.SSS a Z"

	formatted := cleardate.Format(now, customLayout)
	fmt.Printf("Current time formatted: %s\n", formatted)

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Fatalf("Error loading location: %v", err)
	}

	timeStr := "2025-02-09 14:30:00.123 Z"
	customLayout2 := "yyyy-MM-dd HH:mm:ss.SSS Z"
	parsedInLoc, err := cleardate.ParseInLocation(customLayout2, timeStr, loc)
	if err != nil {
		log.Fatalf("Error parsing time in location: %v", err)
	}
	fmt.Printf("Parsed time in %v: %v\n", loc, parsedInLoc)
}
```

## License

Licensed under [MIT License](./LICENSE.md)