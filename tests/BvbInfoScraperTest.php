<?php
namespace pdt256\vbscraper;

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

        // First and last
        $this->verifyMatch($matches[0], 2718, 'Joe Cash', 9024, 'Jonathan Alvarez', '21-9', '21-14', '');
        $this->verifyMatch($matches[74], 7335, 'Brad Keenan', 190, 'Jake Gibb', '16-21', '21-19', '18-16');
        $this->assertSame(75, count($matches));

        // Three set match
        $this->verifyMatch($matches[2], 11097, 'Mike Claman', 1924, 'Matt Heath', '16-21', '21-10', '15-11');
    }

    public function testGetMatchesWithForfeit()
    {
        $content = file_get_contents(__DIR__ . '/2015ManhattanTournament.html');
        $matches = BvbInfoScraper::getMatches($content);

        // First and last
        $this->verifyMatch($matches[0], 11745, 'Connor Hughes', 269, 'Kevin Cleary', '21-7', '21-12', '');
        $this->verifyMatch($matches[102], 5214, 'Phil Dalhausser', 7831, 'Theo Brunner', '21-17', '21-13', '');
        $this->assertSame(103, count($matches));

        // Around Forfeit multi-line issue
        $this->verifyMatch($matches[40], 6908, 'Paul Araiza', 11236, 'Michael Boag', '21-19', '19-21', '19-17');
        $this->verifyMatch($matches[41], 10332, 'Adam Cabbage', 2457, 'Matt Heagy', '', '', '');
        $this->verifyMatch($matches[42], 8143, 'Hawk Hatcher', 6274, 'Robert deAurora', '21-15', '21-17', '');
        $this->verifyMatch($matches[43], 2035, 'John Moran', 15804, 'Michael Brunsting', '21-18', '21-17', '');
    }

    private function verifyMatch(
        Match $match,
        $teamAPlayerAID,
        $teamAPlayerAName,
        $teamBPlayerAID,
        $teamBPlayerAName,
        $score1,
        $score2,
        $score3
    ) {
        $this->assertSame($teamAPlayerAID, $match->getTeamA()->getPlayerA()->getVbId());
        $this->assertSame($teamAPlayerAName, $match->getTeamA()->getPlayerA()->getName());
        $this->assertSame($teamBPlayerAID, $match->getTeamB()->getPlayerA()->getVbId());
        $this->assertSame($teamBPlayerAName, $match->getTeamB()->getPlayerA()->getName());

        if ($score1 !== 'forfeit') {
            $this->assertSame($score1, (string) $match->getSetScore(1));
            $this->assertSame($score2, (string) $match->getSetScore(2));
            $this->assertSame($score3, (string) $match->getSetScore(3));
        }
    }
}
