<?php
namespace pdt256\vbscraper\Service\Import;

use pdt256\vbscraper\Service\BvbInfoScraper;
use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\Entity\SetScore;
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
        $validMatch = $this->getDummyMatch();
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

        $this->setCountLogger();

        $importResult = $this->matchImporter->import($matches);

        $this->assertTrue($importResult instanceof MatchImportResult);
        $this->assertSame(103, $importResult->getSuccessCount());
        $this->assertSame(0, $importResult->getFailedCount());
        $this->assertSame(309, $this->countSQLLogger->getTotalQueries());
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
