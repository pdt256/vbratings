<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\Mapping\ClassMetadata;
use Symfony\Component\Validator\Constraints as Assert;

class Team
{
    /** @var Player */
    private $playerA;

    /** @var Player */
    private $playerB;

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
}
