<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Match;

class MatchRepository extends AbstractEntityRepository implements MatchInterface
{
    public function create(Match & $user)
    {
        $this->createEntity($user);
    }

    public function update(Match & $user)
    {
        $this->updateEntity($user);
    }

    public function delete(Match & $user)
    {
        $this->deleteEntity($user);
    }
}
