<?php
namespace pdt256\vbscraper\tests\Helper\EntityRepository;

use pdt256\vbscraper\EntityRepository\PlayerRepositoryInterface;
use pdt256\vbscraper\Entity\Player;

class FakePlayer extends AbstractFake implements PlayerRepositoryInterface
{
    protected $entityName = 'Player';

    /** @var Player[] */
    protected $entities = [];

    public function findOneByVbId($vbId)
    {
        foreach ($this->entities as $entity) {
            if ($entity->getVbId() === $vbId) {
                return $entity;
            }
        }
    }
}
