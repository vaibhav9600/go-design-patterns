package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client          *http.Client
	timeout         time.Duration
	userAgent       string
	followRedirects bool
}

type Option func(*Client)

func NewClient(opts ...Option) *Client {
	client := Client{
		client:          http.DefaultClient,
		timeout:         30 * time.Second,
		userAgent:       "My Http Client",
		followRedirects: true,
	}

	for _, opt := range opts {
		opt(&client)
	}

	return &client
}

func withTimeout(t time.Duration) Option {
	return func(c *Client) {
		c.timeout = t
	}
}

func withUserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

func withoutRedirect() Option {
	return func(c *Client) {
		c.followRedirects = false
	}
}

func useInsecureTransport() Option {
	return func(c *Client) {
		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
}

func (c *Client) Get(url string) (string, error) {
	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set the User-Agent header if specified
	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	// Perform the HTTP request
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	// Check for a non-200 status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

func main() {
	client := NewClient(
		withTimeout(10*time.Second),
		withUserAgent("My Custom User Agent"),
		useInsecureTransport(),
	)

	// Use the client to make HTTP requests.
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Make the GET request using the custom client
	body, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}

	// Print the response body
	fmt.Println("Response body:")
	fmt.Println(body)
}
