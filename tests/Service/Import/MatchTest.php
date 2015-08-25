<?php
namespace pdt256\vbscraper\Service\Import;

class MatchTest extends \PHPUnit_Framework_TestCase
{
    public function testImport()
    {
        $matchImporter = new Match;

        $importResult = $matchImporter->import();

        $this->assertTrue($importResult instanceof ImportResult);
        $this->assertSame(1, $importResult->getSuccessCount());
        $this->assertSame(1, $importResult->getFailedCount());
    }
}
