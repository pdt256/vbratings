# Volleyball Ratings GraphQL API

## Help

```
$ go run cmd/vbratings-graphql/main.go --help
Volleyball Ratings GraphQL
Usage:
  -dbPath string
        sqlite db path (default "./_data/vb.db")
  -port int
        port (default 8080)
```

## Start GraphQL API Server

```
$ go run cmd/vbratings-graphql/main.go
Volleyball Ratings GraphQL
Starting on port 8080
```

## Example

### Query

```
query($year: Int!, $gender: String!, $limit: Int!) {
  playerRatingQueries {
    getTopPlayerRatings(year: $year, gender: $gender, limit: $limit) {
      player {
        Name
      }
      playerRating {
        Rating
        TotalMatches
      }
    }
  }
}
```

### Variables
```
{
  "year": 2018,
  "gender": "male",
  "limit": 2
}
```

### Response

```
$ curl -s XPOST -d '{"query": "query($year: Int!, $gender: String!, $limit: Int!) { playerRatingQueries { getTopPlayerRatings(year: $year, gender: $gender, limit: $limit) { player { Name } playerRating { Rating TotalMatches } } } }", "variables": {"year": 2018, "gender": "male", "limit": 2} }' localhost:8080/query | python -m json.tool
{
    "data": {
        "playerRatingQueries": {
            "getTopPlayerRatings": [
                {
                    "player": {
                        "Name": "Nick Lucena"
                    },
                    "playerRating": {
                        "Rating": 1921,
                        "TotalMatches": 841
                    }
                },
                {
                    "player": {
                        "Name": "Phil Dalhausser"
                    },
                    "playerRating": {
                        "Rating": 1894,
                        "TotalMatches": 885
                    }
                }
            ]
        }
    }
}
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
