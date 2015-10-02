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

    public function persist(Player & $entity)
    {
        $this->persistEntity($entity);
    }

    public function find($id)
    {
        return $this->findOneBy(['id' => $id]);
    }

    public function findOneByVbId($vbId)
    {
        return $this->findOneBy(['vbId' => $vbId]);
    }
}
