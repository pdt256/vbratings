<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\Team;

interface TeamInterface
{
    public function create(Team & $entity);
    public function update(Team & $entity);
    public function delete(Team & $entity);
    public function persist(Team & $team);

    /**
     * @param int $id
     * @return Team
     */
    public function find($id);

    /**
     * @param Player $playerA
     * @param Player $playerB
     * @return Team|null
     */
    public function findOneByPlayers($playerA, $playerB);
}
