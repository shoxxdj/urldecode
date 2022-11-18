package main

import (
	"fmt"
	"log"
	"net/url"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		encodedValue := scanner.Text()
		decodedValue, err := url.QueryUnescape(encodedValue)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(decodedValue)
	}
}
