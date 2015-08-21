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

        $this->verifyMatch($matches[0], 2718, 'Joe Cash', 9024, 'Jonathan Alvarez', '21-9', '21-14', '');
        $this->verifyMatch($matches[2], 11097, 'Mike Claman', 1924, 'Matt Heath', '16-21', '21-10', '15-11');
        $this->verifyMatch($matches[74], 7335, 'Brad Keenan', 190, 'Jake Gibb', '16-21', '21-19', '18-16');

        $this->assertSame(75, count($matches));
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
        $this->assertSame($score1, (string) $match->getSetScore(1));
        $this->assertSame($score2, (string) $match->getSetScore(2));
        $this->assertSame($score3, (string) $match->getSetScore(3));
    }
}
