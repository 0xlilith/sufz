package main

import (
	"sufz/sufz"
	"flag"
	"bufio"
	"os"
)

func main()  {
	// FLAGS libray
	urlencode := flag.Bool("e", false, "Url Encode")
	urldecode := flag.Bool("d", false, "Url Decode")

	flag.Parse()
	scan := bufio.NewScanner(os.Stdin)

	if *urlencode {
		sufz.UrlEncode(*scan)
	}
	if *urldecode {
		sufz.UrlDecode(*scan)
	}
}