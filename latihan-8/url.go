package main

import (
	"fmt"
	"net/url"
)

func main() {
	uri := "https://www.abdulrasyid.my.id/search?q=golang&level=basic&level=medium"

	url, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	fmt.Println("path", url.Path)
	fmt.Println("scheme", url.Scheme)
	fmt.Println("host", url.Host)
	fmt.Println("Query", url.Query())

}
