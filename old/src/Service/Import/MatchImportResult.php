<?php
namespace pdt256\vbscraper\Service\Import;

use pdt256\vbscraper\Entity;

class MatchImportResult extends AbstractImportResult
{
    /** @var Entity\Match[] */
    protected $failedRows = [];

    public function getFailedRows()
    {
        return $this->failedRows;
    }

    public function addFailedRow(Entity\Match $failedRow)
    {
        $this->failedRows[] = $failedRow;
    }
}
