<?php
namespace pdt256\vbscraper;

class SetScore
{
    /** @var int */
    private $teamAScore;

    /** @var int */
    private $teamBScore;

    /** @var bool */
    private $isTeamBForfeit;

    public function __construct($scoreAsString)
    {
        if ($scoreAsString === 'forfeit') {
            $this->setIsTeamBForfeit(true);
        } else {
            list($teamAScore, $teamBScore) = explode('-', $scoreAsString);
            $this->setTeamAScore($teamAScore);
            $this->setTeamBScore($teamBScore);
        }
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

    public function isTeamBForfeit()
    {
        return $this->isTeamBForfeit;
    }

    /**
     * @param bool $isTeamBForfeit
     */
    public function setIsTeamBForfeit($isTeamBForfeit)
    {
        $this->isTeamBForfeit = (bool) $isTeamBForfeit;
    }

    public function __toString()
    {
        if ($this->isTeamBForfeit()) {
            return 'forfeit';
        } else {
            return $this->getTeamAScore() . '-' . $this->getTeamBScore();
        }
    }
}
