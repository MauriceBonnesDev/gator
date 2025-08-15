package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "gator")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	content := RSSFeed{}
	err = xml.Unmarshal(body, &content)
	if err != nil {
		return nil, err
	}

	content.Channel.Title = html.UnescapeString(content.Channel.Title)
	content.Channel.Description = html.UnescapeString(content.Channel.Description)
	for i, item := range content.Channel.Item {
		content.Channel.Item[i].Title = html.UnescapeString(item.Title)
		content.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &content, nil
}
