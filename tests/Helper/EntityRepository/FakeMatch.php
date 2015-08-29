<?php
namespace pdt256\vbscraper\tests\Helper\EntityRepository;

use pdt256\vbscraper\EntityRepository\MatchInterface;
use pdt256\vbscraper\Entity;

class FakeMatch extends AbstractFake implements MatchInterface
{
    public function __construct()
    {
        $this->setReturnValue(new Entity\Match);
    }

    public function find($id)
    {
        return $this->getReturnValue();
    }

    public function create(Entity\Match & $user)
    {
    }

    public function update(Entity\Match & $user)
    {
    }

    public function delete(Entity\Match & $user)
    {
    }
}
