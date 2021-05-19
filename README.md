# Diretta.it Scraper
A simple web scraper written in Go for m.diretta.it, an italian livescore website.

This project is for test and educational purpose only, I just wanted to try some tools of Golang.

## Prerequisites
- [Docker Compose CE](https://docs.docker.com/compose/install/)

## Installation
```bash
docker-compose exec direttascraper-golang bash -c "make setup"
```

## Run
```bash
docker-compose exec direttascraper-golang bash -c "make run"
```
An Echo server will start on port `17171`, so you can send requests to `localhost:17171`.

### Debug
If you want to start the Delve debugger:
```bash
docker-compose exec direttascraper-golang bash -c "make debug"
```

## API
### GET /v1/matches
Returns a list of daily matches
```json
[
  {
    "name": "Talleres (R.E) - Temperley",
    "date": "19.05.2021 20:10",
    "status": "ENDED",
    "result": "0:4"
  },
  {
    "name": "Ind. Juniors - Cumbaya",
    "date": "19.05.2021 22:00",
    "status": "LIVE",
    "result": "0:0"
  },
  {
    "name": "Inter Nouakchott - Armee",
    "date": "19.05.2021 21:40",
    "status": "SCHEDULED"
  },
  {
    "name": "Visakha (Cam) - Lalenok (Tls)",
    "date": "19.05.2021 12:00",
    "status": "POSTPONED"
  },
  {
    ...
  }
]
```