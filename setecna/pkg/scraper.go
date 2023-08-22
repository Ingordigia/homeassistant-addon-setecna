package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

type HTMLMeta struct {
	CsrfToken string `json:"csrf-token"`
}

// Login
func Login(client *http.Client) {
	// Parse URL
	u, err := url.Parse(idrosistemiLoginURL)
	if err != nil {
		log.Fatal(err)
	}

	// Recover token
	resp, err := client.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	meta := extract(resp.Body)

	log.Printf(meta.CsrfToken)

	data := url.Values{}
	data.Set("_token", meta.CsrfToken)
	data.Set("email", user)
	data.Set("password", password)

	_, err = client.PostForm(u.String(), data)
	if err != nil {
		log.Fatal(err)
	}
}

func AskRefresh(client *http.Client) {
	u, err := url.Parse(idrosistemiRefreshURL)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}

func Fetch(client *http.Client, oldtimestamp string) (IdrosistemiResponse, string) {

	var requestURL string
	var newTimestamp string

	start := time.Now()

	if oldtimestamp != "" {
		requestURL = idrosistemiUpdateURL + oldtimestamp
	} else {
		requestURL = idrosistemiUpdateURL
	}

	u, err := url.Parse(requestURL)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var result IdrosistemiResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// // Build a config map:
	// confMap := map[string]int{}
	// for _, v := range result.Data {
	// 	confMap[v.ID] = v.V
	// }
	newTimestamp = result.Timestamp

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
	return result, newTimestamp
}

func Push(client *http.Client, sensor string, value string) (err error) {
	fullURL := idrosistemiChangeURL + "&p0=" + sensor + "&nb0=" + value
	u, err := url.Parse(fullURL)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(resp.Body)

	defer resp.Body.Close()
	return nil
}

func extract(resp io.Reader) *HTMLMeta {
	z := html.NewTokenizer(resp)

	hm := new(HTMLMeta)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return hm
			}
			if t.Data == "meta" {
				csrfToken, ok := extractMetaProperty(t, "csrf-token")
				if ok {
					hm.CsrfToken = csrfToken
				}
			}
		}
	}
}

func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if attr.Key == "name" && attr.Val == prop {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
		}
	}

	return
}

func Scrape() {

	// Create Cookie Jar
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	// Create client
	client := &http.Client{
		Jar: jar,
	}

}
