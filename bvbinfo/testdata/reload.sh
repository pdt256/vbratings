#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

curl -s "http://bvbinfo.com/season.asp" -o "$DIR/all-seasons.html"
curl -s "http://bvbinfo.com/Season.asp?AssocID=1&Year=2017" -o "$DIR/2017-avp-tournaments.html"
curl -s "http://bvbinfo.com/Tournament.asp?ID=3332&Process=Matches" -o "$DIR/2017-avp-manhattan-beach-mens-matches.html"
curl -s "http://www.bvbinfo.com/Tournament.asp?ID=3333&Process=Matches" -o "$DIR/2017-avp-manhattan-beach-womens-matches.html"
curl -s "http://bvbinfo.com/Tournament.asp?ID=3109&Process=Matches" -o "$DIR/2015-avp-manhattan-beach-mens-matches.html"
curl -s "http://bvbinfo.com/Tournament.asp?ID=2975&Process=Matches" -o "$DIR/2014-avp-st-petersburg-mens-matches.html"
curl -s "http://www.bvbinfo.com/Tournament.asp?ID=3465&Process=Matches" -o "$DIR/2018-fivb-gstaad-major-mens-matches.html"
curl -s "http://bvbinfo.com/player.asp?ID=1256" -o "$DIR/misty-may-player.html"
