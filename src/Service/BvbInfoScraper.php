<?php
namespace pdt256\vbscraper\Service;

use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\SetScore;
use pdt256\vbscraper\Entity\Team;

class BvbInfoScraper
{
    /**
     * @param string $xmlContent
     * @return string[]
     */
    public static function getSeasonUrls($xmlContent)
    {
        $xpath = '//a[contains(@href, "Season")]/@href';
        $nodes = static::getDomNodeList($xmlContent, $xpath);

        $urls = [];
        foreach ($nodes as $node) {
            $urls[] = $node->value;
        }

        return $urls;
    }

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
     * @return \pdt256\vbscraper\Entity\Match[]
     */
    public static function getMatches($xmlContent)
    {
        $regex = '/
            \<br\>Match\s\d+:
            [^?]+
            \?ID=(?<team1PlayerAID>\d+)"\>(?<team1PlayerAName>[^<]+)\<\/a\>
            [^?]+
            \?ID=(?<team1PlayerBID>\d+)"\>(?<team1PlayerBName>[^<]+)\<\/a\>
            [^?]+
            \?ID=(?<team2PlayerAID>\d+)"\>(?<team2PlayerAName>[^<]+)\<\/a\>
            [^?]+
            \?ID=(?<team2PlayerBID>\d+)"\>(?<team2PlayerBName>[^<]+)\<\/a\>
            [^?]+
            \)
            (?:
                (?:
                    \sby\s(?<forfeit>Forfeit)
                )
                |
                (?:
                    \s(?<score1>\d+-\d+)
                    ,\s(?<score2>\d+-\d+)
                    (,\s(?<score3>\d+-\d+))?
                    \s\((?<time>\d+:\d+)\)
                )
            )
        /xm';

        preg_match_all($regex, $xmlContent, $regexMatches, PREG_SET_ORDER);

        $matches = [];

        foreach ($regexMatches as $regexMatch) {
            $teamA = new Team;
            $teamA->setPlayerA(new Player(
                $regexMatch['team1PlayerAID'],
                $regexMatch['team1PlayerAName']
            ));
            $teamA->setPlayerB(new Player(
                $regexMatch['team1PlayerBID'],
                $regexMatch['team1PlayerBName']
            ));

            $teamB = new Team;
            $teamB->setPlayerA(new Player(
                $regexMatch['team2PlayerAID'],
                $regexMatch['team2PlayerAName']
            ));
            $teamB->setPlayerB(new Player(
                $regexMatch['team2PlayerBID'],
                $regexMatch['team2PlayerBName']
            ));

            $match = new Match;
            $match->setTeamA($teamA);
            $match->setTeamB($teamB);

            if ($regexMatch['forfeit'] === 'Forfeit') {
                $match->addSetScore(new SetScore('forfeit'));
            } else {
                $match->addSetScore(new SetScore($regexMatch['score1']));
                $match->addSetScore(new SetScore($regexMatch['score2']));

                $score3 = $regexMatch['score3'];
                if ($score3 !== '') {
                    $match->addSetScore(new SetScore($score3));
                }
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
}
