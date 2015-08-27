<?php
namespace pdt256\vbscraper\Service;

use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\Entity\Player;

class BvbInfoScraperTest extends \PHPUnit_Framework_TestCase
{
    public function testGetSeasonUrls()
    {
        $content = file_get_contents(__DIR__ . '/AllSeasons.html');
        $tournamentUrls = BvbInfoScraper::getSeasonUrls($content);

        $this->assertSame('Season.asp?AssocID=1&Year=2015', $tournamentUrls[0]);
        $this->assertSame('Season.asp?AssocID=4&Year=1948', $tournamentUrls[240]);
        $this->assertSame(241, count($tournamentUrls));
    }

    public function testGetSeasonTournamentUrls()
    {
        $content = file_get_contents(__DIR__ . '/2014SeasonTournaments.html');
        $tournamentUrls = BvbInfoScraper::getSeasonTournamentUrls($content);

        $this->assertSame('Tournament.asp?ID=2975', $tournamentUrls[0]);
        $this->assertSame('Tournament.asp?ID=2988', $tournamentUrls[13]);
        $this->assertSame(14, count($tournamentUrls));
    }

    public function testGetMatches()
    {
        $content = file_get_contents(__DIR__ . '/2014StPTournament.html');
        $matches = BvbInfoScraper::getMatches($content);

        $this->assertSame(75, count($matches));

        $this->verifyMatch(
            $matches[0],
            new Player(2718, 'Joe Cash'), new Player(10141, 'Roger Corbitt'),
            new Player(9024, 'Jonathan Alvarez'), new Player(13516, 'Shaun Dawson'),
            '21-9', '21-14', ''
        );

        $this->verifyMatch(
            $matches[74],
            new Player(7335, 'Brad Keenan'), new Player(5323, 'John Mayer'),
            new Player(190, 'Jake Gibb'), new Player(5327, 'Casey Patterson'),
            '16-21', '21-19', '18-16'
        );
    }

    public function testGetMatchesWithForfeit()
    {
        $content = file_get_contents(__DIR__ . '/2015ManhattanTournament.html');
        $matches = BvbInfoScraper::getMatches($content);

        $this->assertSame(103, count($matches));

        $this->verifyMatch(
            $matches[40],
            new Player(6908, 'Paul Araiza'), new Player(10333, 'Alejandro Parra'),
            new Player(11236, 'Michael Boag'), new Player(6277, 'Dan Newman'),
            '21-19', '19-21', '19-17'
        );
        $this->verifyMatch(
            $matches[41],
            new Player(10332, 'Adam Cabbage'), new Player(13455, 'Jake Rosener'),
            new Player(2457, 'Matt Heagy'), new Player(13522, 'Jorge Martinez'),
            'forfeit', '', ''
        );
        $this->verifyMatch(
            $matches[42],
            new Player(8143, 'Hawk Hatcher'), new Player(14828, 'Clay Paullin'),
            new Player(6274, 'Robert deAurora'), new Player(10199, 'Ozz Borges'),
            '21-15', '21-17', ''
        );
        $this->verifyMatch(
            $matches[43],
            new Player(2035, 'John Moran'), new Player(4825, 'Gregg Weaver'),
            new Player(15804, 'Michael Brunsting'), new Player(9301, 'Hylas Smith'),
            '21-18', '21-17', ''
        );
    }

    private function verifyMatch(
        Match $match,
        Player $teamAPlayerA,
        Player $teamAPlayerB,
        Player $teamBPlayerA,
        Player $teamBPlayerB,
        $score1,
        $score2,
        $score3
    ) {
        $this->assertEquals($teamAPlayerA, $match->getTeamA()->getPlayerA());
        $this->assertEquals($teamAPlayerB, $match->getTeamA()->getPlayerB());
        $this->assertEquals($teamBPlayerA, $match->getTeamB()->getPlayerA());
        $this->assertEquals($teamBPlayerB, $match->getTeamB()->getPlayerB());
        $this->assertSame($score1, (string) $match->getSetScore(1));
        $this->assertSame($score2, (string) $match->getSetScore(2));
        $this->assertSame($score3, (string) $match->getSetScore(3));
    }
}
