package main

import (
	"fmt"
	"math/rand"
	"time"

	"ts/nwo/gen/attributes"
)

func main() {

	randy := rand.New(rand.NewSource(time.Now().UnixNano()))
	atts := attributes.RandomAtttributes(func() attributes.Level {
		raw := randy.Int() % 5 //[0,5]
		return attributes.ToLevelRoundUp(2 * (raw + 2))
	})

	fmt.Println("| Attribute | Level |")
	fmt.Println("|:----|----:|")
	for _, attr := range atts {
		fmt.Printf("| %s | %s |\n", attr.Name(), attr.Level().String())
	}
}
