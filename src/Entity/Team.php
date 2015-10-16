<?php
namespace pdt256\vbscraper\Entity;

use Doctrine\Common\Collections\ArrayCollection;
use Symfony\Component\Validator\Mapping\ClassMetadata;
use Symfony\Component\Validator\Constraints as Assert;

class Team implements EntityInterface, ValidationInterface
{
    use IdTrait, TimeTrait;

    /** @var Player */
    private $playerA;

    /** @var Player */
    private $playerB;

    /** @var Match[] */
    private $matches;

    public function __construct()
    {
        $this->setCreated();
        $this->matches = new ArrayCollection;
    }

    public static function loadValidatorMetadata(ClassMetadata $metadata)
    {
        $metadata->addPropertyConstraint('playerA', new Assert\Valid);
        $metadata->addPropertyConstraint('playerB', new Assert\Valid);
    }

    public function getPlayerA()
    {
        return $this->playerA;
    }

    public function setPlayerA(Player $playerA)
    {
        $this->playerA = $playerA;
    }

    public function getPlayerB()
    {
        return $this->playerB;
    }

    public function setPlayerB(Player $playerB)
    {
        $this->playerB = $playerB;
    }

    public function __toString()
    {
        return '[' . (string) $this->getPlayerA() . ', ' . (string) $this->getPlayerB() . ']';
    }
}
