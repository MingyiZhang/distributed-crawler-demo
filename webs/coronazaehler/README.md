# coronazaehler.de crawler

It scrapes the current COVID-19 data of Germany in [coronazaehler.de](coronazaehler.de) and
saves to a elasticsearch node county by county.

## How to run

Run everything: 
```shell script
docker-compose up --build
```
Open `localhost:8888` in a web browser. 
Search with county's name or state's name. 

## Clean up
stop everything:
```shell script
docker-compose down
```