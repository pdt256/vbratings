<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Match;

class MatchRepository extends AbstractEntityRepository implements MatchInterface
{
    public function create(Match & $entity)
    {
        $this->createEntity($entity);
    }

    public function update(Match & $entity)
    {
        $this->updateEntity($entity);
    }

    public function delete(Match & $entity)
    {
        $this->deleteEntity($entity);
    }
}
