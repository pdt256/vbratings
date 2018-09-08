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

#### Initialize Database

```
$ go run bvbinfo/cmd/import/main.go -init
BVBInfo Importer
Initializing DB
```

#### Import from BvbInfo

```
$ go run bvbinfo/cmd/import/main.go
BVBInfo Importer
Importing Matches
...............................................................................
...............................................................................
...............................................................................
...............................................................................
109531 matches imported
Importing Players
...............................................................................
...............................................................................
260 players imported
```

---

### Volleyball Ratings

```
$ go run cmd/vbratings/main.go --help
Volleyball Ratings
Usage:
  -dbPath="./_data/vb.db": sqlite db path
  -init=false: init db
  -year=2018: year
```

### Calculate ratings for 2018 (seeded from previous year if exists)

```
$ go run cmd/vbratings/main.go -year 2018
BVBInfo Ratings
1846 ratings calculated
```

## Todo

* Use Cases:
  - As a fan of the sport, I want to see all player ratings by gender and year; so that
    I can view the top players over time.
  - As a player, I want to see my rating for each year; so that I can monitor
    my progress relative to other players.
* Tasks:
  - Add # of games played to player_rating table and increment when matches are calculated

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
