# Distributed Web Crawler Demo

A simple distributed web crawler library that is written in Go. 

The library is implemented completed from scratch, as a practice of Go. 
It is the capstone project of the imooc's Golang [course](https://coding.imooc.com/class/180.html).

## Architecture
As a distributed web crawler, it contains several components
- crawler engine
- persistent service
- worker for parsing web data
Components are communicated through JSON-RPC.

## Algorithm



## Examples
There are two simple examples included:
- [Coronazaehler](./webs/coronazaehler):
scrape current COVID-19 data of every county in Germany from [coronazaehler.de](https://www.coronazaehler.de/).
- [mockweb](./webs/mockweb): scrape profile data from a mock dating website.

## TODO
- [x] separate service for saving data
- [x] separate service for parsing web data
- [x] frontend for display search results
- [ ] separate service for checking duplication
- [ ] Kubernetes deployment

 