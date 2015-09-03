<?php
namespace pdt256\vbscraper\tests\Helper;

use Doctrine\DBAL\Logging\SQLLogger;

class CountSQLLogger implements SQLLogger
{
    /** @var int */
    protected $totalQueries;

    public function getTotalQueries()
    {
        return $this->totalQueries;
    }

    public function __construct()
    {
        $this->totalQueries = 0;
    }

    public function startQuery($sql, array $params = null, array $types = null)
    {
        //$this->displaySql($sql, $params);

        if (! $this->isTransactionSQL($sql)) {
            $this->totalQueries++;
        }
    }

    private function displaySql($sql, $params = null)
    {
        $values = '';
        if (is_array($params)) {
            $values = json_encode(array_values($params));
        }

        echo $sql . ' ' . $values . PHP_EOL . PHP_EOL;
    }

    public function stopQuery()
    {
    }

    /**
     * @param string $sql
     * @return bool
     */
    private function isTransactionSQL($sql)
    {
        return substr($sql, 0, 1) === '"';
    }
}
