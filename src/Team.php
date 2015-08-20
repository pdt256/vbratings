<?php
namespace pdt256\vbscraper;

class Team
{
    /** @var Player */
    private $playerA;

    /** @var Player */
    private $playerB;

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
