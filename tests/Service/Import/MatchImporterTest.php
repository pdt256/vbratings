<?php
namespace pdt256\vbscraper\Service\Import;

use pdt256\vbscraper\Service\BvbInfoScraper;
use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\Entity\Team;
use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\SetScore;
use pdt256\vbscraper\tests\Helper\EntityRepository\FakeMatch;
use pdt256\vbscraper\tests\Helper;

class MatchImporterTest extends Helper\DoctrineTestCase
{
    protected $metaDataClassNames = [
        'vbscraper:Match',
    ];

    /** @var MatchImporter */
    protected $matchImporter;

    public function setUp()
    {
        $this->matchImporter = new MatchImporter($this->repository()->getMatch());
    }

    public function testImportWithValidAndInvalid()
    {
        $validMatch = $this->getValidMatch();
        $invalidMatch = $this->getInvalidMatch();
        $matches = [
            $validMatch,
            $invalidMatch
        ];

        $this->setCountLogger();

        $importResult = $this->matchImporter->import($matches);

        $failedRows = $importResult->getFailedRows();

        $this->assertTrue($importResult instanceof MatchImportResult);
        $this->assertSame(1, $importResult->getSuccessCount());
        $this->assertSame(1, $importResult->getFailedCount());
        $this->assertSame(1, count($failedRows));
        $this->assertSame(1, count($importResult->getErrorMessages()));
        $this->assertSame($invalidMatch, $failedRows[0]);
        $this->assertSame(3, $this->countSQLLogger->getTotalQueries());
    }

    public function testImportFromTournament()
    {
        $content = file_get_contents(__DIR__ . '/../2015ManhattanTournament.html');
        $matches = BvbInfoScraper::getMatches($content);

        $importResult = $this->matchImporter->import($matches);

        $this->assertTrue($importResult instanceof MatchImportResult);
        $this->assertSame(103, $importResult->getSuccessCount());
        $this->assertSame(0, $importResult->getFailedCount());
    }

    private function getValidMatch()
    {
        $teamA = new Team();
        $teamA->setPlayerA(new Player(1, 'John Doe'));
        $teamA->setPlayerB(new Player(2, 'James Doe'));

        $teamB = new Team();
        $teamB->setPlayerA(new Player(3, 'John Smith'));
        $teamB->setPlayerB(new Player(4, 'James Smith'));

        $setScore = new SetScore;
        $setScore->setScoresByString('21-18');

        $match = new Match;
        $match->setTeamA($teamA);
        $match->setTeamB($teamB);
        $match->addSetScore($setScore);

        return $match;
    }

    private function getInvalidMatch()
    {
        $setScore = new SetScore;
        $setScore->setTeamAScore(-1);
        $setScore->setTeamBScore(-1);

        $match = new Match;
        $match->addSetScore($setScore);

        return $match;
    }
}
