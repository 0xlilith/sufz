package sufz

import (
	"bufio"
	"fmt"
	"net/url"
)

func UrlEncode(scan bufio.Scanner) {
	for scan.Scan() {
		fmt.Println(url.QueryEscape(scan.Text()))
	}
}