<?php
namespace pdt256\vbscraper\Entity;

use RuntimeException;
use Symfony\Component\Validator\Mapping\ClassMetadata;
use Symfony\Component\Validator\Constraints as Assert;

class Match implements EntityInterface
{
    use IdTrait, TimeTrait;

    /** @var Team */
    private $teamA;

    /** @var Team */
    private $teamB;

    /** @var SetScore[] */
    private $setScores;

    /** @var int */
    private $gameTimeLengthInSeconds;

    public function __construct()
    {
        $this->setScores = [];
        $this->setCreated();
    }

    public static function loadValidatorMetadata(ClassMetadata $metadata)
    {
        $metadata->addPropertyConstraint('teamA', new Assert\Valid);
        $metadata->addPropertyConstraint('teamB', new Assert\Valid);
        $metadata->addPropertyConstraint('setScores', new Assert\Valid);
        $metadata->addPropertyConstraint('gameTimeLengthInSeconds', new Assert\Range([
            'min' => 0,
            'max' => 65535,
        ]));
    }

    public function getTeamA()
    {
        return $this->teamA;
    }

    public function setTeamA(Team $teamA)
    {
        $this->teamA = $teamA;
    }

    public function getTeamB()
    {
        return $this->teamB;
    }

    public function setTeamB(Team $teamB)
    {
        $this->teamB = $teamB;
    }

    public function getSetScores()
    {
        return $this->setScores;
    }

    public function addSetScore(SetScore $setScore)
    {
        if (count($this->setScores) === 3) {
            throw new RuntimeException('Exceeded maximum set scores.');
        }

        $this->setScores[] = $setScore;
    }

    public function addSetScoreByString($scoreAsString)
    {
        $setScore = new SetScore;
        $setScore->setScoresByString($scoreAsString);
        $this->addSetScore($setScore);
    }

    public function getGameTimeLengthInSeconds()
    {
        return $this->gameTimeLengthInSeconds;
    }

    /**
     * @param int $gameTimeLengthInSeconds
     */
    public function setGameTimeLengthInSeconds($gameTimeLengthInSeconds)
    {
        $this->gameTimeLengthInSeconds = (int) $gameTimeLengthInSeconds;
    }

    /**
     * @param int $set
     * @return mixed
     */
    public function getSetScore($set)
    {
        $index = (int) $set - 1;
        if (! isset($this->setScores[$index])) {
            return null;
        }

        return $this->setScores[$index];
    }

    public function __toString()
    {
        return (string) $this->getTeamA() . ' def ' . (string) $this->getTeamB();
    }
}
