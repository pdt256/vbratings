<?php
namespace pdt256\vbscraper;

class BvbInfoScraper
{
    /**
     * @param string $xmlContent
     * @return string[]
     */
    public static function getSeasonTournamentUrls($xmlContent)
    {
        $xpath = '//a[contains(@href, "Tournament")]/@href';
        $nodes = static::getDomNodeList($xmlContent, $xpath);

        $urls = [];

        foreach ($nodes as $node) {
            $urls[] = $node->value;
        }

        return $urls;
    }

    /**
     * @param string $xmlContent
     * @return Match[]
     */
    public static function getMatches($xmlContent)
    {
        $regex = '/
            Match\s\d*:
            .*
            ID=(?<team1PlayerAID>\d*)"\>(?<team1PlayerAName>.*)\<\/a\>
            .*
            ID=(?<team1PlayerBID>\d*)"\>(?<team1PlayerBName>.*)\<\/a\>
            .*
            ID=(?<team2PlayerAID>\d*)"\>(?<team2PlayerAName>.*)\<\/a\>
            .*
            ID=(?<team2PlayerBID>\d*)"\>(?<team2PlayerBName>.*)\<\/a\>
            .*
            \)
            \s(?<score1>\d+-\d+)
            ,\s(?<score2>\d+-\d+)
            (,\s(?<score3>\d+-\d+))?
            \s\((?<time>\d*:\d*)\)
        /x';

        preg_match_all($regex, $xmlContent, $regexMatches);

        $matches = [];

        $total = count($regexMatches['team1PlayerAID']);
        for ($i = 0; $i < $total; $i++) {
            $teamA = new Team;
            $teamA->setPlayerA(new Player(
                $regexMatches['team1PlayerAID'][$i],
                $regexMatches['team1PlayerAName'][$i]
            ));
            $teamA->setPlayerB(new Player(
                $regexMatches['team1PlayerBID'][$i],
                $regexMatches['team1PlayerBName'][$i]
            ));

            $teamB = new Team;
            $teamB->setPlayerA(new Player(
                $regexMatches['team2PlayerAID'][$i],
                $regexMatches['team2PlayerAName'][$i]
            ));
            $teamB->setPlayerB(new Player(
                $regexMatches['team2PlayerBID'][$i],
                $regexMatches['team2PlayerBName'][$i]
            ));

            $match = new Match;
            $match->setTeamA($teamA);
            $match->setTeamB($teamB);

            $match->addSetScore(new SetScore($regexMatches['score1'][$i]));
            $match->addSetScore(new SetScore($regexMatches['score2'][$i]));

            $score3 = $regexMatches['score3'][$i];
            if ($score3 !== '') {
                $match->addSetScore(new SetScore($score3));
            }

            $matches[] = $match;
        }

        return $matches;
    }

    /**
     * @param string $xmlContent
     * @param string $xpath
     * @return \DOMNodeList
     */
    private static function getDomNodeList(& $xmlContent, & $xpath)
    {
        $domDocument = new \DOMDocument();
        libxml_use_internal_errors(true);
        $domDocument->loadHTML($xmlContent);

        $domXpath = new \DOMXPath($domDocument);
        return $domXpath->query($xpath);
    }

    /**
     * @param $matches
     */
    private static function viewMatchData($matches, $index)
    {
        print_r([
            $matches['team1PlayerAID'][$index],
            $matches['team1PlayerAName'][$index],
            $matches['team1PlayerBID'][$index],
            $matches['team1PlayerBName'][$index],
            $matches['team2PlayerAID'][$index],
            $matches['team2PlayerAName'][$index],
            $matches['team2PlayerBID'][$index],
            $matches['team2PlayerBName'][$index],
            $matches['score1'][$index],
            $matches['score2'][$index],
            $matches['score3'][$index],
            $matches['time'][$index],
        ]);
    }
}
