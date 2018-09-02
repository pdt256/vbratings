Volleyball Score Scraper
========================
[![Build Status](https://travis-ci.org/pdt256/vbscraper.svg?branch=master)](https://travis-ci.org/pdt256/vbscraper)

Download volleyball match results

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

### CLI Application

```
$ go run cmd/cli/bvbinfo/main.go --help
BVBInfo Importer
Usage of bvbinfo:
  -allSeasons=false: load all seasons
  -dbPath="./_data/vb.db": sqlite db path
  -init=false: init db
  -seasonUrl="": season url
  -tournamentUrl="": tournament url
```

#### Initialize Database

```
$ go run cmd/cli/bvbinfo/main.go -init
```

#### Import Matches from a Tournament

```
$ go run cmd/cli/bvbinfo/main.go -tournamentUrl "http://bvbinfo.com/Tournament.asp?ID=3320&Process=Matches"
BVBInfo Importer
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3320&Process=Matches
84 matches imported
$ sqlite3 _data/vb.db 'select * from match limit 3;'
4007e6ea-1e98-4310-8aa0-7185809e2e0a|1171|11005|17456|7060
e6d2cd6b-9bde-40e7-bc62-c7dfeed7fada|14846|16729|6274|17246
284e5ec8-cdc2-4e0c-9835-c1f12c0d9da1|6908|7376|16023|6276
```

#### Import Matches from a Season

```
$ go run cmd/cli/bvbinfo/main.go -seasonUrl "http://bvbinfo.com/Season.asp?AssocID=1&Year=2017"
BVBInfo Importer
Importing Season: http://bvbinfo.com/Season.asp?AssocID=1&Year=2017
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3320&Process=Matches
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3321&Process=Matches
...
1389 matches imported
$ sqlite3 _data/vb.db 'select count(*) from match;'
1389
```

#### Import Matches from all Seasons

```
$ go run cmd/cli/bvbinfo/main.go -allSeasons
BVBInfo Importer
Importing Season: http://bvbinfo.com/Season.asp?AssocID=3&Year=2019
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3547&Process=Matches
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3548&Process=Matches
...
Importing Season: http://bvbinfo.com/Season.asp?AssocID=1&Year=2018
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3485&Process=Matches
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3486&Process=Matches
...
109531 matches imported
```

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
