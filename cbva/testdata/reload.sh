#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

getTournament() {
    id=$1
    file=$2

    curl -s 'https://cbva.com/Results/GetTournamentTeamResult' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' --data-binary "{\"id\":\"${id}\"}" --compressed > "${DIR}/${file}"
}

getTournament A14CC0CB1B90719A "2018-09-23-marine-street-mens-aa.json"
getTournament D718CFD486CD14D6 "2018-09-01-hermosa-mens-open.json"
