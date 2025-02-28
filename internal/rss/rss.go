package rss

// import (
// 	"context"
// 	"encoding/xml"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// )

// type RSSFeed struct {
// 	Channel struct {
// 		Title       string    `xml:"title"`
// 		Link        string    `xml:"link"`
// 		Description string    `xml:"description"`
// 		Item        []RSSItem `xml:"item"`
// 	} `xml:"channel"`
// }

// type RSSItem struct {
// 	Title       string `xml:"title"`
// 	Link        string `xml:"link"`
// 	Description string `xml:"description"`
// 	PubDate     string `xml:"pubDate"`
// }

// func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
// 	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating request: %w", err)
// 	}

// 	req.Header.Set("User-Agent", "gator")

// 	client := http.Client{Timeout: 5 * time.Second}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating request: %w", err)
// 	}

// 	defer resp.Body.Close()

// 	rssfeed := &RSSFeed{}
// 	data, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("error reading response body %w", err)
// 	}

// 	err = xml.Unmarshal(data, rssfeed)
// 	if err != nil {
// 		return nil, fmt.Errorf("error decoding xml %w", err)
// 	}
// 	return rssfeed, nil
// }
