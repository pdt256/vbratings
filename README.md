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
$ go run cmd/cli/bvbinfo/main.go -init
```

```
$ go run cmd/cli/bvbinfo/main.go -tournamentUrl "http://bvbinfo.com/Tournament.asp?ID=3332&Process=Matches"
$ sqlite3 _data/vb.db 'select * from match limit 3;'
9888df7f-cfad-4a51-8749-c7a8269deec7|7376|2037|17858|17468
26741cf2-ac74-41c7-8e83-96a69ce44b2f|11745|15817|17083|18012
0987d6d2-057c-4627-8a92-d986619e4165|11237|18008|18033|18032
```

```
$ go run cmd/cli/bvbinfo/main.go -seasonUrl "http://bvbinfo.com/Season.asp?AssocID=1&Year=2017"
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3320&Process=Matches (84) matches imported
Importing Tournament: http://bvbinfo.com/Tournament.asp?ID=3321&Process=Matches (76) matches imported
...
Done! (1389) matches imported
$ sqlite3 _data/vb.db 'select count(*) from match;'
1389
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
