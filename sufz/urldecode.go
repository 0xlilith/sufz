package sufz

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
)

func UrlDecode(scan bufio.Scanner) {
	for scan.Scan() {
		value, err := url.QueryUnescape(scan.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	}
}