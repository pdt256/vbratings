<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;

class PlayerRepository extends AbstractEntityRepository implements PlayerInterface
{
    public function create(Player & $entity)
    {
        $this->createEntity($entity);
    }

    public function update(Player & $entity)
    {
        $this->updateEntity($entity);
    }

    public function delete(Player & $entity)
    {
        $this->deleteEntity($entity);
    }

    public function findOneByVbId($vbId)
    {
        return $this->findOneBy(['vbId' => $vbId]);
    }
}