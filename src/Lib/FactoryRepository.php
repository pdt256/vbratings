<?php
namespace pdt256\vbscraper\Lib;

use Doctrine\ORM\EntityManager;
use pdt256\vbscraper\EntityRepository;

class FactoryRepository
{
    /** @var EntityManager */
    protected $entityManager;

    public function __construct(EntityManager $entityManager)
    {
        $this->entityManager = $entityManager;
    }

    /**
     * @return EntityRepository\MatchInterface
     */
    public function getMatch()
    {
        return $this->entityManager->getRepository('vbscraper:Match');
    }

    /**
     * @return EntityRepository\TeamInterface
     */
    public function getTeam()
    {
        return $this->entityManager->getRepository('vbscraper:Team');
    }

    /**
     * @return EntityRepository\PlayerInterface
     */
    public function getPlayer()
    {
        return $this->entityManager->getRepository('vbscraper:Player');
    }
}
