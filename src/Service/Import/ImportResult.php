<?php
namespace pdt256\vbscraper\Service\Import;

class ImportResult
{
    /** @var int */
    private $successCount = 0;

    /** @var array[] */
    private $failedRows = [];

    /** @var string[] */
    private $errorMessages = [];

    public function getSuccessCount()
    {
        return $this->successCount;
    }

    public function getFailedCount()
    {
        return count($this->failedRows);
    }

    public function incrementSuccess()
    {
        $this->successCount++;
    }

    public function getFailedRows()
    {
        return $this->failedRows;
    }

    public function addFailedRow(array $failedRow)
    {
        $this->failedRows[] = $failedRow;
    }

    public function addErrorMessage($errorMessage)
    {
        $this->errorMessages[] = $errorMessage;
    }

    public function getErrorMessages()
    {
        return $this->errorMessages;
    }
}
