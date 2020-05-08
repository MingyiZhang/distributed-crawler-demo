# Distributed Web Crawler Demo

A simple distributed web crawler library that is written in Go. 

The web crawler is implemented completed from scratch, as a practice of Go. 

## Architecture
As a distributed web crawler, it contains 
- crawler engine
- persistent service
- worker for parsing web 


## Examples
There are two simple examples included:
- [Coronazaehler](./webs/coronazaehler):
scrape current COVID-19 data of every county in Germany from [coronazaehler.de](https://www.coronazaehler.de/).
- [mockweb](./webs/mockweb): scrape profile data from a mock dating website.

## TODO
- [ ] separate service for checking duplication
- [ ] Kubernetes deployment

 