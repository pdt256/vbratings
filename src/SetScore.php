<?php
namespace pdt256\vbscraper;

class SetScore
{
    /** @var int */
    private $teamAScore;

    /** @var int */
    private $teamBScore;

    public function __construct($scoreAsString)
    {
        list($teamAScore, $teamBScore) = explode('-', $scoreAsString);
        $this->setTeamAScore($teamAScore);
        $this->setTeamBScore($teamBScore);
    }

    public function getTeamAScore()
    {
        return $this->teamAScore;
    }

    /**
     * @param int $teamAScore
     */
    public function setTeamAScore($teamAScore)
    {
        $this->teamAScore = (int) $teamAScore;
    }

    public function getTeamBScore()
    {
        return $this->teamBScore;
    }

    /**
     * @param int $teamBScore
     */
    public function setTeamBScore($teamBScore)
    {
        $this->teamBScore = (int) $teamBScore;
    }
}
