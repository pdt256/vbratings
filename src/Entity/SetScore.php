<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\GroupSequenceProviderInterface;
use Symfony\Component\Validator\Mapping\ClassMetadata;
use Symfony\Component\Validator\Constraints as Assert;

class SetScore implements GroupSequenceProviderInterface
{
    /** @var int */
    private $teamAScore;

    /** @var int */
    private $teamBScore;

    /** @var bool */
    private $teamBForfeit;

    public function __construct()
    {
        $this->unsetTeamBForfeit();
    }

    public static function loadValidatorMetadata(ClassMetadata $metadata)
    {
        $metadata->setGroupSequenceProvider(true);

        $metadata->addPropertyConstraint('teamAScore', new Assert\NotNull);
        $metadata->addPropertyConstraint('teamAScore', new Assert\Range([
            'min' => 0,
            'max' => 64,
        ]));

        $metadata->addPropertyConstraint('teamBScore', new Assert\NotNull);
        $metadata->addPropertyConstraint('teamBScore', new Assert\Range([
            'min' => 0,
            'max' => 64,
        ]));

        $metadata->addPropertyConstraint('teamBForfeit', new Assert\NotNull([
            'groups' => ['SetScore', 'Forfeit'],
        ]));
    }

    public function getGroupSequence()
    {
        $groups = ['SetScore'];

        if ($this->isTeamBForfeit()) {
            $groups = ['Forfeit'];
        }

        return $groups;
    }

    public function setScoresByString($scoreAsString)
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

    public function isTeamBForfeit()
    {
        return $this->teamBForfeit;
    }

    public function setTeamBForfeit()
    {
        $this->teamBForfeit = true;
    }

    public function unsetTeamBForfeit()
    {
        $this->teamBForfeit = false;
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
