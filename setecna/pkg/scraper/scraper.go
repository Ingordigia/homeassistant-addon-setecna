package scraper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Ingordigia/homeassistant-addon-setecna/models"
	"github.com/Ingordigia/homeassistant-addon-setecna/pkg/mqtt"
	broker "github.com/eclipse/paho.mqtt.golang"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

const baseURL = "https://s5a.eu"

type htmlMeta struct {
	CsrfToken string `json:"csrf-token"`
}

func extract(resp io.Reader) *htmlMeta {
	z := html.NewTokenizer(resp)

	hm := new(htmlMeta)

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

type FlexString string

func (fi *FlexString) UnmarshalJSON(b []byte) error {
	if b[0] == '"' {
		return json.Unmarshal(b, (*string)(fi))
	}
	var i int
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	s := strconv.Itoa(i)
	*fi = FlexString(s)
	return nil
}

type Response struct {
	Status    string `json:"Status"`
	Timestamp string `json:"Timestamp"`
	Latest    string `json:"Latest"`
	Rows      int    `json:"Rows"`
	Data      []struct {
		ID string     `json:"Id"`
		V  FlexString `json:"V"`
	} `json:"Data"`
}

func (r *Response) GetUpdatedValues(systemID string, params models.ParamsMap) (msgs []mqtt.Message) {
	var message mqtt.Message
	for _, sensor := range r.Data {
		if _, ok := params[sensor.ID]; ok {
			if sensor.ID == "LAST_UPDATE" {
				if sensor.V != "" {
					date, err := time.Parse("2006-01-02 15:04:05.000000-07", string(sensor.V))
					if err != nil {
						log.Println(err, ", ignoring")
						continue
					} else {
						message = mqtt.Message{
							Topic:   "homeassistant/sensor/" + systemID + "_LAST_UPDATE",
							Message: date.Format(time.RFC3339),
							Qos:     0,
						}
					}
				} else {
					continue
				}
			} else {
				message = mqtt.Message{
					Topic:   "homeassistant/" + params[sensor.ID].EntityType + "/" + systemID + "_" + sensor.ID,
					Message: string(sensor.V),
					Qos:     0,
				}
			}
			msgs = append(msgs, message)
		}
	}
	return msgs
}

type Scraper struct {
	client          *http.Client
	loginURL        string
	fetchUpdatesURL string
	askRefreshURL   string
	pushUpdatesURL  string
	lastFetch       string
}

func (s *Scraper) push(key string, value string) (err error) {
	fullURL := s.pushUpdatesURL + "&p0=" + key + "&nb0=" + value
	u, err := url.Parse(fullURL)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := s.client.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	return nil
}

func (s *Scraper) Init(systemID string) {
	// Create Cookie Jar
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	// Create client
	c := &http.Client{
		Jar: jar,
	}
	s.client = c
	s.loginURL = baseURL + "/login"
	s.fetchUpdatesURL = baseURL + "/station/" + systemID + "/getres?timestamp="
	s.askRefreshURL = baseURL + "/station/" + systemID + "/askrefresh?connrq=1"
	s.pushUpdatesURL = baseURL + "/station/" + systemID + "/putmprop?statid=" + systemID + "&userid=guest&pcount=1" //&p0=ACS_SET_COMFORT&nb0=400
}

func (s *Scraper) Login(user, password string) error {
	// Parse URL
	u, err := url.Parse(s.loginURL)
	if err != nil {
		return err
	}

	// Recover token
	resp, err := s.client.Get(u.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	meta := extract(resp.Body)

	// log.Printf(meta.CsrfToken)

	data := url.Values{}
	data.Set("_token", meta.CsrfToken)
	data.Set("email", user)
	data.Set("password", password)

	_, err = s.client.PostForm(u.String(), data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scraper) AskRefresh() error {
	u, err := url.Parse(s.askRefreshURL)
	if err != nil {
		return err
	}

	resp, err := s.client.Get(u.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

// func (s *Scraper) Fetch(oldtimestamp string) (Response, string, string) {

// 	var requestURL, newTimestamp, lastUpdate string

// 	if oldtimestamp != "" {
// 		requestURL = s.fetchUpdatesURL + url.QueryEscape(oldtimestamp)
// 	} else {
// 		requestURL = s.fetchUpdatesURL
// 	}

// 	u, err := url.Parse(requestURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	resp, err := s.client.Get(u.String())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	var result Response
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		fmt.Println("Can not unmarshal JSON", err)
// 	}

// 	fmt.Println("Fetch Response: ", result)
// 	newTimestamp = result.Timestamp
// 	lastUpdate = result.Latest

// 	return result, newTimestamp, lastUpdate
// }

func (s *Scraper) Fetch() (Response, error) {

	var requestURL string

	if s.lastFetch != "" {
		requestURL = s.fetchUpdatesURL + url.QueryEscape(s.lastFetch)
	} else {
		requestURL = s.fetchUpdatesURL
	}

	u, err := url.Parse(requestURL)
	if err != nil {
		return Response{}, err
	}

	resp, err := s.client.Get(u.String())
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON", err)
	}

	s.lastFetch = result.Timestamp

	result.Data = append(result.Data, struct {
		ID string     "json:\"Id\""
		V  FlexString "json:\"V\""
	}{
		ID: "LAST_UPDATE",
		V:  FlexString(result.Latest),
	})

	return result, nil
}

func (s *Scraper) Change(systemID string) broker.MessageHandler {
	return func(client broker.Client, msg broker.Message) {
		key := strings.Split(msg.Topic(), "/")[2]
		if strings.HasPrefix(key, systemID) {
			key = strings.TrimPrefix(key, systemID+"_")
			log.Printf("Changing %s to %s\n", key, string(msg.Payload()))
			s.push(key, string(msg.Payload()))
		} else {
			log.Printf("Received a message for another systemID: %s, ignoring\n", systemID)
		}

	}
}
