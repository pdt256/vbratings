<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;

interface PlayerInterface
{
    public function create(Player & $entity);
    public function update(Player & $entity);
    public function delete(Player & $entity);
    public function persist(Player & $player);

    /**
     * @param int $id
     * @return Player
     */
    public function find($id);

    /**
     * @param int $vbId
     * @return Player|null
     */
    public function findOneByVbId($vbId);
}
