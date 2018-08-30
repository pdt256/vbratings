<?php
namespace pdt256\vbscraper\Service\Import;

abstract class AbstractImportResult
{
    protected $failedRows = [];

    /** @var int */
    private $successCount = 0;

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

    /**
     * @param string $errorMessage
     */
    public function addErrorMessage($errorMessage)
    {
        $this->errorMessages[] = (string) $errorMessage;
    }

    public function getErrorMessages()
    {
        return $this->errorMessages;
    }
}
