<?php
namespace pdt256\vbscraper\EntityRepository;

use Doctrine;
use pdt256\vbscraper\Entity\EntityInterface;

abstract class AbstractEntityRepository extends Doctrine\ORM\EntityRepository
{
    public function createEntity(EntityInterface & $entity)
    {
        $entityManager = $this->getEntityManager();
        $entityManager->persist($entity);
        $entityManager->flush();
    }

    public function updateEntity(EntityInterface & $entity)
    {
        $entityManager = $this->getEntityManager();
        $entity = $entityManager->merge($entity);
        $entityManager->flush();
    }

    public function deleteEntity(EntityInterface $entity)
    {
        $entityManager = $this->getEntityManager();
        $entityManager->remove($entity);
        $entityManager->flush();
    }
}
