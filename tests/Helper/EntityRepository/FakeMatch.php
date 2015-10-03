<?php
namespace pdt256\vbscraper\tests\Helper\EntityRepository;

use pdt256\vbscraper\EntityRepository\MatchInterface;
use pdt256\vbscraper\Entity\Match;

class FakeMatch extends AbstractFake implements MatchInterface
{
    public function __construct()
    {
        $this->setReturnValue(new Match);
    }

    public function find($id)
    {
        return $this->getReturnValue();
    }

    public function create(Match & $user)
    {
    }

    public function update(Match & $user)
    {
    }

    public function delete(Match & $user)
    {
    }

    public function persist(Match & $match)
    {
    }

    public function clear()
    {
    }
}
