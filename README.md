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

## Run

### Bvbinfo Importer

```
$ go run bvbinfo/cmd/import/main.go -h
BVBInfo Importer
Usage:
  -dbPath="./_data/vb.db": sqlite db path
  -init=false: init db
```

#### Import

```
$ go run bvbinfo/cmd/import/main.go -init
BVBInfo Importer
Initializing DB
```

```
$ go run bvbinfo/cmd/import/main.go
BVBInfo Importer
Importing Matches
...............................................................................
...............................................................................
...............................................................................
94218 matches imported
Importing Players
...............................................................................
...............................................................................
...............................................................................
...............................................................................
11373 players imported
```

---

### Calculate Volleyball Ratings

```
$ go run cmd/calculate-vbratings/main.go --help
Volleyball Ratings Calculator
Usage:
  -allYears=false: calculate for all years
  -dbPath="./_data/vb.db": sqlite db path
  -init=false: init db
  -year=2018: year
```

```
$ go run cmd/calculate-vbratings/main.go -init
Volleyball Ratings Calculator
Initializing player_rating DB
```

#### Calculate ratings for all years

```
$ go run cmd/calculate-vbratings/main.go -allYears
Volleyball Ratings Calculator
..................
20532 ratings calculated
```

---

### Volleyball Ratings

```
$ go run cmd/view-vbratings/main.go -topPlayers -gender male -year 2018 -limit 10
Volleyball Ratings
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

---

## Todo

* Use Cases:
  - As a player, I want to see my rating for each year; so that I can monitor
    my progress relative to other players.

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
