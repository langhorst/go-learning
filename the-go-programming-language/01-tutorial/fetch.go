// Fetch prints the content found at each specified URL.
//
// Exercise 1.7 uses `io.Copy(dst, src)` to read from `src` and write to `dst`
// instead of using `ioutil.ReadAll` to copy the response body to `os.Stdout`
// and requiring a buffer large enough to hold the entire stream.
//
// Exercise 1.8 adds an `http://` prefix to each argument if it's not already
// present.
//
// Exercise 1.9 also prints the HTTP status code.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix("http://", url) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, resp.Status+"\n")
	}
}
