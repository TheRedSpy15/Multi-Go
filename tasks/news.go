package tasks

import (
	"fmt"

	"github.com/SlyMarbo/rss"
	"github.com/TheRedSpy15/Multi-Go/utils"
)

// News uses the ycombinator rss feed to print the front page titles
func News() {
	feed, err := rss.Fetch("https://news.ycombinator.com/rss")
	if err != nil {
		utils.CheckErr(err)
	}

	for _, item := range feed.Items {
		fmt.Println("Title :", item.Title)
	}
}
