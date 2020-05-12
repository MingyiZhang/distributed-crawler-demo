# Simple Distributed Web Crawler Library

A simple distributed web crawler library that is written in Go. 

The library is implemented completed from scratch. 
As a Golang practice project, it is mainly focused on the distributed structure. 
One needs to implement their own web parsers as shown in the [examples](#examples).

It is the capstone project of the imooc's Golang [course](https://coding.imooc.com/class/180.html).

## Architecture
As a distributed web crawler, it contains several components
- [Concurrent engine](./engine/concurrent.go) manages the crawler's core logic among components.
    - [Queued Scheduler](./scheduler/queued.go) manages workers and requests in queues.
- [Persistent service](./persist) is for saving scraped data. Right now it saves parsed data into elasticsearch. More database can be supported.
- [Crawler worker service](./worker) is for parsing website. 

Components are communicated using JSON-RPC.

## Algorithm
The crawler uses breadth first search to scrape website.

## Examples
There are two simple examples included:
- [Coronazaehler](./webs/coronazaehler) scrapes current COVID-19 data of every county in Germany from [coronazaehler.de](https://www.coronazaehler.de/).
- [mockweb](./webs/mockweb) scrapes profile data from a mock dating website.

## Note
As shown in the 
## TODO
- [x] separate service for saving data
- [x] separate service for parsing web data
- [x] frontend for display search results
- [ ] separate service for checking duplication
- [ ] Kubernetes deployment
- [ ] gRPC and Protobuf version
 