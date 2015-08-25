<?php
namespace pdt256\vbscraper\Service\Import;

class Match
{
    public function import()
    {
        $importResult = new ImportResult;
        $importResult->incrementSuccess();
        $importResult->addFailedRow([]);

        return $importResult;
    }
}
