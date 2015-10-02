<?php
namespace pdt256\vbscraper\tests\Helper\EntityRepository;

use pdt256\vbscraper\EntityRepository\PlayerInterface;
use pdt256\vbscraper\Entity\Player;

class FakePlayer extends AbstractFake implements PlayerInterface
{
    public function __construct()
    {
        $this->setReturnValue(new Player);
    }

    public function create(Player & $user)
    {
    }

    public function update(Player & $user)
    {
    }

    public function delete(Player & $user)
    {
    }

    public function persist(Player & $player)
    {
    }

    public function find($id)
    {
        return $this->getReturnValue();
    }

    public function findOneByVbId($vbId)
    {
        return $this->getReturnValue();
    }
}
