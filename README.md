# go-scrape

## Introduction
A simple web crawler, written in Go. The crawler is limited to one domain (e.g. if you start with https://www.bbc.co.uk/, it would crawl all pages within bbc.co.uk, but not follow external links, for example to the Facebook and Twitter accounts.)

Given a URL, it will print a simple site map, showing the links between pages.

## Installation
Before running this Go project you will need to install the following packages:

- `go get github.com/PuerkitoBio/goquery`

## Usage
The crawler accepts two command line flag arguments:

- `startPage` - the URL the crawler should start with
- `timeout` - the timeout of the HTTP GET request to each page (note: if the request fails, it will do so silently.)

Example usage:

The following example will run the crawler, starting with `http://www.bbc.co.uk`, and will give each GET request a `timeout` of 30 seconds.

```
go run crawl.go --startPage http://www.bbc.co.uk --timeout 30
```

You can also run `go run crawl.go --help` to see a list of the command line arguments.

## Future Improvements

- Provide the user with a 'depth' argument, which will allow them to control how many levels deep the scraper will go. 

- Keep a track of any website that fail (e.g. because of a error status code) and log them out at the end so the user is aware.

## Output

While the crawler is running, you will see a sequence of dots (`.`) that indicate a page has been scraped. This is purely to indicate to the user that something is happening. 

Once the crawler has finished, an output will be displayed on the screen that followings the following convention:

```
- URI: {Link} (The page that has been crawled)
|
|--- Linked URL: {Link} - {Count} (The page that was found, and a count of how many times)
|--- Linked URL: {Link} - {Count} (The page that was found, and a count of how many times)
|--- Linked URL: {Link} - {Count} (The page that was found, and a count of how many times)
```

This will be repeated for every page crawled. 