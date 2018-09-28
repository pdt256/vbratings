Volleyball Ratings
========================
[![Build Status](https://travis-ci.org/pdt256/vbratings.svg?branch=master)](https://travis-ci.org/pdt256/vbratings)

Calculate volleyball player ratings from public match results

## Setup

### Install Dependencies

```
go get ./...
```

## Unit Tests

```
go test ./...
```

## Sub Packages

- [BVBInfo Importer](bvbinfo/README.md)
- [CBVA Importer](cbva/README.md)

---

## Calculate Volleyball Ratings

```
$ go run cmd/calculate/main.go --help
Volleyball Ratings Calculator
Usage:
  -allYears
        calculate for all years
  -dbPath string
        sqlite db path (default "./_data/vb.db")
  -init
        init db
  -year int
        year (default 2018)
```

```
$ go run cmd/calculate/main.go -init
Volleyball Ratings Calculator
Initializing player_rating DB
```

```
$ go run cmd/calculate/main.go -allYears
Volleyball Ratings Calculator
..................
20532 ratings calculated
```

---

## Volleyball Ratings

### CLI

```
$ go run cmd/cli/main.go --help
Usage:
  app [command]

Available Commands:
  help        Help about any command
  topPlayers  List Top Players By Year

Flags:
  -d, --dbPath string   sqlite db path (default "./_data/vb.db")
  -h, --help            help for app

Use "app [command] --help" for more information about a command.
```

```
$ go run cmd/cli/main.go topPlayers -h
List Top Players By Year

Usage:
  app topPlayers  [flags]

Flags:
  -g, --gender string   gender (default "male")
  -h, --help            help for topPlayers
  -l, --limit int       limit (default 10)
  -y, --year int        year (default 2018)
```

```
$ go run cmd/cli/main.go topPlayers --gender male --year 2018 --limit 10
Top 10 male Players in 2018
+--------+-------------------------+--------------+
| RATING |          NAME           | TOTALMATCHES |
+--------+-------------------------+--------------+
|   1921 | Nick Lucena             |          841 |
|   1894 | Phil Dalhausser         |          885 |
|   1892 | Alexander Brouwer       |          359 |
|   1892 | Robert Meeuwsen         |          338 |
|   1874 | Anders Berntsen Mol     |           90 |
|   1856 | Pablo Herrera Allepuz   |          429 |
|   1856 | Adrián Gavira Collado   |          407 |
|   1853 | Paolo Nicolai           |          308 |
|   1853 | Daniele Lupo            |          282 |
|   1850 | Christian Sandlie Sørum |          131 |
+--------+-------------------------+--------------+
```

### GraphQL

```
$ go run cmd/vbratings-graphql/main.go --help
Volleyball Ratings GraphQL
Usage:
  -dbPath string
        sqlite db path (default "./_data/vb.db")
  -port int
        port (default 8080)
```

#### Start GraphQL API Server

```
$ go run cmd/vbratings-graphql/main.go
Volleyball Ratings GraphQL
Starting on port 8080
```

#### Example Query

```
{
  topPlayers(year: 2018, gender: "male", limit: 2) {
    rating
    playerName
    totalMatches
  }
}
```

#### Example Response

```
$ curl -s -XPOST -d '{"query": "{ getTopPlayerRatings(year: 2018, gender: \"male\", limit: 2) { rating playerName totalMatches } }"}' localhost:8080/query | python -m json.tool
{
    "data": {
        "getTopPlayerRatings": [
            {
                "playerName": "Nick Lucena",
                "rating": 1921,
                "totalMatches": 841
            },
            {
                "playerName": "Phil Dalhausser",
                "rating": 1894,
                "totalMatches": 885
            }
        ]
    }
}
```

---

## Todo

* Use Cases:
  - As a player, I want to see my rating for each year; so that I can monitor
    my progress relative to other players.
  - As a fan of the sport, I want to see all player ratings by gender and year that
    include results from CBVA, AVPFirst, and AVPNext; while not colliding with AVP results
    from Bvbinfo; so that I can see more accurate player rankings among lower
    rated players.
    - Import results from CBVA
    - Normalize player ID to work across tournament organizations by using full name slugs
    - Calculate rankings based on tournament results in addition to match results.

## License

The MIT License (MIT)

Copyright (c) 2015 Jamie Isaacs <pdt256@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
