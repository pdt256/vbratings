package vbscraper

import (
	"io"
	"io/ioutil"
	"regexp"
)

type match struct {
	PlayerA player
	PlayerB player
	PlayerC player
	PlayerD player
}

type player struct {
	Id   string
	Name string
}

func GetMatches(matchReader io.Reader) []match {
	playerRegex := `<a href="player.asp\?ID=(\d+)">([\w\s]+)</a>`
	re := regexp.MustCompile(playerRegex + `.*?` + playerRegex + `.*?` + playerRegex + `.*?` + playerRegex)
	bytes, _ := ioutil.ReadAll(matchReader)
	regexMatches := re.FindAllStringSubmatch(string(bytes), -1)

	var matches []match
	for _, value := range regexMatches {
		matches = append(matches, match{
			player{value[1], value[2]},
			player{value[3], value[4]},
			player{value[5], value[6]},
			player{value[7], value[8]},
		})
	}

	return matches
}
