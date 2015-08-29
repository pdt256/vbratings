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
     * @param EntityManager $entityManager
     * @return self
     */
    public static function getInstance(EntityManager $entityManager)
    {
        static $factoryRepository = null;

        if ($factoryRepository === null) {
            $factoryRepository = new static($entityManager);
        }

        return $factoryRepository;
    }

    /**
     * @return EntityRepository\MatchInterface
     */
    public function getMatch()
    {
        return $this->entityManager->getRepository('vbscraper:Match');
    }
}
