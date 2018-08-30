<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\Team;

/**
 * @method Team findOneById($id)
 */
interface TeamRepositoryInterface extends AbstractRepositoryInterface
{
    /**
     * @param Player $playerA
     * @param Player $playerB
     * @return Team|null
     */
    public function findOneByPlayers($playerA, $playerB);
}
