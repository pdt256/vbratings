<?php
namespace pdt256\vbscraper;

class BvbInfoScraperTest extends \PHPUnit_Framework_TestCase
{
    public function testGetSeaonTournamentUrls()
    {
        $content = file_get_contents(__DIR__ . '/2014SeasonTournaments.html');
        $tournamentUrls = BvbInfoScraper::getSeasonTournamentUrls($content);

        $count = 0;
        foreach ($tournamentUrls as $node) {
            $count++;
        }

        $this->assertSame('Tournament.asp?ID=2975', $tournamentUrls[0]->value);
        $this->assertSame('Tournament.asp?ID=2988', $tournamentUrls[13]->value);
        $this->assertSame(14, $count);
    }

    public function testGetMatchResults()
    {
        $content = file_get_contents(__DIR__ . '/2014StPTournament.html');
        $matches = BvbInfoScraper::getMatches($content);

        $this->assertSame(2718, $matches[0]->getTeamA()->getPlayerA()->getVbId());
        $this->assertSame('Joe Cash', $matches[0]->getTeamA()->getPlayerA()->getName());
        $this->assertSame(9024, $matches[0]->getTeamB()->getPlayerA()->getVbId());
        $this->assertSame('Jonathan Alvarez', $matches[0]->getTeamB()->getPlayerA()->getName());
        $this->assertSame(75, count($matches));
    }
}
