package api

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Client is a Valr API client.
type Client struct {
	httpClient   *http.Client
	baseURL      string
	apiKeyPub    string
	apiKeySecret string
	debug        bool
}

const defaultBaseURL = "https://api.valr.com/v1"

const defaultTimeout = 10 * time.Second

// NewClient creates a new Valr API client with the default base URL.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: defaultTimeout},
		baseURL:    defaultBaseURL,
	}
}

// SetAuth provides the client with an API key and secret.
func (cl *Client) SetAuth(apiKeyID, apiKeySecret string) error {
	if apiKeyID == "" || apiKeySecret == "" {
		return errors.New("valr: no credentials provided")
	}
	cl.apiKeyPub = apiKeyID
	cl.apiKeySecret = apiKeySecret
	return nil
}

// SetHTTPClient sets the HTTP client that will be used for API calls.
func (cl *Client) SetHTTPClient(httpClient *http.Client) {
	cl.httpClient = httpClient
}

// SetTimeout sets the timeout for requests made by this client. Note: if you
// set a timeout and then call .SetHTTPClient(), the timeout in the new HTTP
// client will be used.
func (cl *Client) SetTimeout(timeout *time.Duration) {
	if timeout == nil {
		cl.httpClient.Timeout = defaultTimeout
	} else {
		cl.httpClient.Timeout = *timeout
	}
}

// SetBaseURL overrides the default base URL. For internal use.
func (cl *Client) SetBaseURL(baseURL string) {
	cl.baseURL = strings.TrimRight(baseURL, "/")
}

// SetDebug enables or disables debug mode. In debug mode, HTTP requests and
// responses will be logged.
func (cl *Client) SetDebug(debug bool) {
	cl.debug = debug
}

func (cl *Client) do(ctx context.Context, method, path string,
	req, res interface{}, auth bool) error {

	url := cl.baseURL + "/" + strings.TrimLeft(path, "/")

	if cl.debug {
		log.Printf("valr: Call: %s %s", method, path)
		log.Printf("valr: Request: %#v", req)
	}

	var body []byte
	if req != nil {
		values, err := MakeURLValues(req)
		if err != nil {
			return err
		}
		tags := findTags(path)
		for _, tag := range tags {
			url = strings.Replace(url, "{"+tag+"}", values.Get(tag), -1)
			values.Del(tag)
		}
		for key := range values {
			if values.Get(key) == "" {
				values.Del(key)
			}
		}
		if method == http.MethodGet {
			if values.Encode() != "" {
				url = url + "?" + values.Encode()
			}
		} else {
			body, err = json.Marshal(req)
			if err != nil {
				return err
			}
		}
	}

	httpReq, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	httpReq = httpReq.WithContext(ctx)

	if method != http.MethodGet {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	if auth {
		httpReq.Header.Set("X-VALR-API-KEY", cl.apiKeyPub)
		now := time.Now()
		timestampString := strconv.FormatInt(now.UnixNano()/1000000, 10)
		path := strings.Replace(url, "https://api.valr.com", "", -1)
		signature := signRequest(cl.apiKeySecret, timestampString, method, path, body)
		httpReq.Header.Set("X-VALR-SIGNATURE", signature)
		httpReq.Header.Set("X-VALR-TIMESTAMP", timestampString)
	}

	httpRes, err := cl.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpRes.Body.Close()

	body, err = ioutil.ReadAll(httpRes.Body)
	if err != nil {
		return err
	}
	if cl.debug {
		log.Printf("Response: %s", string(body))
	}

	if httpRes.StatusCode/100 != 2 {
		return fmt.Errorf("valr: error response (%d %s)",
			httpRes.StatusCode, http.StatusText(httpRes.StatusCode))
	}

	return json.Unmarshal(body, res)
}

func findTags(str string) []string {
	tags := make([]string, 0)
	firstTag, remainder := findTag(str)
	if firstTag != "" {
		tags = append(tags, firstTag)
	}
	for remainder != "" {
		var nextTag string
		nextTag, remainder = findTag(remainder)
		if nextTag != "" {
			tags = append(tags, nextTag)
		}
	}
	return tags
}

func findTag(str string) (string, string) {
	leftIndex := strings.Index(str, "{")
	if leftIndex > 0 {
		rightIndex := strings.Index(str[leftIndex:], "}")
		if rightIndex >= 0 {
			return str[leftIndex+1 : rightIndex+leftIndex], str[rightIndex+leftIndex:]
		}
	}
	return "", ""
}

func signRequest(apiSecret string, timestampString, verb, path string, body []byte) string {
	// Create a new Keyed-Hash Message Authentication Code (HMAC) using SHA512 and API Secret
	mac := hmac.New(sha512.New, []byte(apiSecret))

	mac.Write([]byte(timestampString))
	mac.Write([]byte(strings.ToUpper(verb)))
	mac.Write([]byte(path))
	mac.Write(body)
	// Gets the byte hash from HMAC and converts it into a hex string
	return hex.EncodeToString(mac.Sum(nil))
}
