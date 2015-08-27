<?php
namespace pdt256\vbscraper\Service\Import;

use pdt256\vbscraper\Service;

class MatchTest extends \PHPUnit_Framework_TestCase
{
    public function testImport()
    {
        $matchImporter = new Match;

        $content = file_get_contents(__DIR__ . '/../2015ManhattanTournament.html');
        $matches = Service\BvbInfoScraper::getMatches($content);

        $importResult = $matchImporter->import($matches);

        $this->assertTrue($importResult instanceof MatchImportResult);
        $this->assertSame(103, $importResult->getSuccessCount());
        $this->assertSame(0, $importResult->getFailedCount());
    }
}
