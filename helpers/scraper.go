package helpers

import (
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

func getSrc(t html.Token) (src string) {
	// loop through tokens attributes
	for _, attr := range t.Attr {
		if attr.Key == "src" {
			src = attr.Val
		}
	}

	// finished search
	return
}

// GetImgsFromURL gets the list of images from a given url
func GetImgsFromURL(url string) (imgs []string, err error) {
	// grab response from link fetch
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	// grab body
	body := resp.Body

	// close after recieved
	defer body.Close()

	// tokenise request
	tokens := html.NewTokenizer(body)

	// loop through tokens
	for {
		// move to next one
		token := tokens.Next()

		// if we have reached the end
		if token == html.ErrorToken {
			return
		}

		// check for img
		if tag := tokens.Token(); tag.Data == "img" {
			// try get url
			src := getSrc(tag)

			// compile regex to match url
			r, err := regexp.Compile(`(?:([^:/?#]+):)?(?://([^/?#]*))?([^?#]*\.(?:jpg|gif|png|jpeg))(?:\?([^#]*))?(?:#(.*))?`)

			// match url element to see if it's a legit image
			if r.MatchString(src) && err == nil {
				// just show dawg
				imgs = append(imgs, src)
			}
		}
	}
}
