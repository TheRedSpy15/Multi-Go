package tasks

/**
Copyright 2018 TheRedSpy15

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/SlyMarbo/rss"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/gocolly/colly"
)

// TODO: need to import from github instead
func News() {
	PrintSources()
	Sources(GetChoice("Select a news source : "))
}

// PrintSources prints all the available news sources and their sourceIDs
func PrintSources() {
	fmt.Println(`
	---- Tech ----
	( 1 ) Dev.to
	( 2 ) Hackernews
	( 3 ) Arstechnica

	---- Politics ----
	( 4 ) New york Times
	( 5 ) Washington Post
	( 6 ) CNN
	( 7 ) Foxnews
	( 8 ) Economist
	( 9 ) USA Today
	( 10 ) Politico

	---- World ----
	( 11 ) theguardian
	( 12 ) npr
	( 13 ) Wall Street Journal

	---- Military ----
	( 14 ) Defense News
	`)
}

// PrintOptions displays available choices after an article has been selected
func PrintOptions() {
	fmt.Println(`
	Select what you want to do with this article
	( 1 ) View - Beta
	( 2 ) Download
	`)
}

// GetChoice gets the index number selected by the user, in order to select it later on
func GetChoice(msg string) int {
	ct.Foreground(ct.Blue, true)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	ct.ResetColor()

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimRight(choice, "\n")
	choiceID, _ := strconv.Atoi(choice)

	return choiceID
}

// News select news source to print articles from, based on ID
// TODO: re-order switch
func Sources(sourceID int) {
	var source string

	// select source based on sourceID
	switch sourceID {
	case 1:
		source = "https://dev.to/feed" // dev.to
	case 2:
		source = "https://news.ycombinator.com/rss" // hackernews
	case 4:
		source = "https://rss.nytimes.com/services/xml/rss/nyt/Politics.xml" // nytimes
	case 5:
		source = "http://feeds.washingtonpost.com/rss/politics" // post
	case 6:
		source = "http://rss.cnn.com/rss/cnn_allpolitics.rss" // cnn
	case 7:
		source = "http://feeds.foxnews.com/foxnews/politics" // fox
	case 3:
		source = "http://feeds.arstechnica.com/arstechnica/index" // ars
	case 13:
		source = "https://www.wsj.com/xml/rss/3_7085.xml" // wsj
	case 8:
		source = "http://www.economist.com/blogs/democracyinamerica/index.xml" // economist
	case 12:
		source = "https://www.npr.org/rss/rss.php?id=1004" // npr
	case 11:
		source = "https://www.theguardian.com/world/rss" // guardian
	case 9:
		source = "http://rssfeeds.usatoday.com/usatodaycomwashington-topstories&x=1" // usa today
	case 10:
		source = "https://www.politico.com/rss/politics08.xml" // politico
	case 14:
		source = "https://feeds.feedburner.com/defense-news/pentagon" // defense
	default:
		fmt.Println("Invalid selection")
		Sources(GetChoice("Select a news source : "))
	}

	// get feed from source
	feed, err := rss.Fetch(source)
	if err != nil {
		panic(err.Error())
	}

	// Select article
	articles := IndexArticles(feed)
	DisplayTitles(articles)
	article := articles[GetChoice("Select an article : ")]

	// Article options
	PrintOptions()
	switch GetChoice("Choice : ") {
	case 1:
		ViewArticle(article) // View
	case 2:
		err := DownloadArticle(article) // Download
		if err != nil {
			panic(err.Error())
		}
	}
}

// DownloadArticle saves an html file copy of the article
func DownloadArticle(article *rss.Item) error {
	name := article.Title

	// Make sure file is easily openable
	if !strings.Contains(name, ".html") {
		name = name + ".html"
	}

	// Create the file
	out, err := os.Create(name)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(article.Link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// DisplayTitles loops through an entire array of rss items and prints the titles
func DisplayTitles(articles []*rss.Item) {
	index := 0
	for _, item := range articles {
		fmt.Println("(", index, ")", item.Title)
		index++
	}
}

// IndexArticles loops through an entire rss feed and returns an array of items
func IndexArticles(feed *rss.Feed) []*rss.Item {
	articles := make([]*rss.Item, len(feed.Items))
	index := 0
	for _, item := range feed.Items {
		articles[index] = item // * using append function causes memory errors strangely
		index++
	}

	return articles
}

// ViewArticle uses colly to scrape an article's link for "p" html tag
// * npr has a cookie policy problem
func ViewArticle(article *rss.Item) {
	c := colly.NewCollector(
		colly.Async(false),
		colly.UserAgent("CollectiveGoColly"),
		colly.IgnoreRobotsTxt(),
	)
	c.DisableCookies()

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("p"))
		fmt.Println(e.Text)
	})

	c.Visit(article.Link)
}
