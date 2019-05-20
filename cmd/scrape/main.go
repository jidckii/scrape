package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jidckii/scrape"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	flag.CommandLine.SetOutput(os.Stdout)

	baseURL := flag.String("url", "https://vedhavyas.com", "Starting URL")
	maxDepth := flag.Int("max-depth", -1, "Max depth to Crawl")
	domainRegex := flag.String("domain-regex", "", "Domain regex to limit crawls to. Defaults to base url domain")
	sitemapFile := flag.String("sitemap", "", "File location to write sitemap to")
	minionCount := flag.Int("minion-count", -1, "Count of parallel minions running. If -1, is count of cpu * 2")
	help := flag.Bool("help", false, "Show Options")
	flag.Parse()

	if *help {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if *baseURL == "" {
		log.Fatal("start URL cannot be empty")
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	resp, err := scrape.StartWithDepthAndDomainRegex(ctx, *baseURL, *maxDepth, *domainRegex, *minionCount)
	if err != nil {
		log.Fatalf("couldn't start scrape: %v\n", err)
	}

	if *sitemapFile != "" {
		scrape.Sitemap(resp, *sitemapFile)
		return
	}

	fmt.Print(resp)
}
