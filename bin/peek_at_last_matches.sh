#!/bin/bash
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

sqlite3 $DIR/../data/db.sqlite -batch -separator $'\t' '
SELECT
    Match.id
    ,GROUP_CONCAT(SetScore.teamAScore || "-" || SetScore.teamBScore)
    ,TAPlayerA.name || "-" || TAPlayerB.name,
    TBPlayerA.name || "-" || TBPlayerB.name
FROM
    Match
    JOIN SetScore ON SetScore.match_id = Match.id
    JOIN Team AS TeamA ON TeamA.id = Match.teamA_id
    Join Team as TeamB ON TeamB.id = Match.teamB_id
    JOIN Player as TAPlayerA ON TAPlayerA.id = TeamA.playerA_id
    JOIN Player as TAPlayerB ON TAPlayerB.id = TeamA.playerB_id
    JOIN Player as TBPlayerA ON TBPlayerA.id = TeamB.playerA_id
    JOIN Player as TBPlayerB ON TBPlayerB.id = TeamB.playerB_id
GROUP BY Match.id
ORDER BY Match.id DESC
LIMIT 20
' | column -ts $'\t'
