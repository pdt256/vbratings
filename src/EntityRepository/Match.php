<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity;

class Match extends AbstractEntityRepository implements MatchInterface
{
    public function create(Entity\Match & $user)
    {
        $this->createEntity($user);
    }

    public function update(Entity\Match & $user)
    {
        $this->updateEntity($user);
    }

    public function delete(Entity\Match & $user)
    {
        $this->deleteEntity($user);
    }
}
