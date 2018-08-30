<?php
namespace pdt256\vbscraper\tests\Helper\EntityRepository;

use pdt256\vbscraper\Entity\EntityInterface;
use pdt256\vbscraper\Entity\ValidationInterface;
use pdt256\vbscraper\EntityRepository\AbstractRepositoryInterface;
use pdt256\vbscraper\EntityRepository\EntityNotFoundException;

class AbstractFake implements AbstractRepositoryInterface
{
    /** @var ValidationInterface[] */
    protected $entities = [];

    protected $entityName = 'Entity';

    /**
     * @param int $id
     * @return EntityInterface
     * @throws EntityNotFoundException
     */
    public function findOneById($id)
    {
        if (isset($this->entities[$id])) {
            return $this->entities[$id];
        }

        throw $this->getEntityNotFoundException();
    }

    protected function getEntityNotFoundException()
    {
        return new EntityNotFoundException($this->entityName . ' not found');
    }

    public function getQueryBuilder()
    {
    }

    public function update(EntityInterface & $entity)
    {
        if (method_exists($entity, 'setUpdated')) {
            $entity->setUpdated();
        }

        $this->entities[$entity->getId()] = $entity;
    }

    public function create(EntityInterface & $entity)
    {
        $entity->setId($this->getAutoincrement());
        $this->entities[$entity->getId()] = $entity;
    }

    public function delete(EntityInterface $entity)
    {
        if (isset($this->entities[$entity->getId()])) {
            unset($this->entities[$entity->getId()]);
        }
    }

    public function persist(EntityInterface & $entity)
    {
    }

    public function merge(EntityInterface & $entity)
    {
    }

    public function flush()
    {
    }

    private function getAutoincrement()
    {
        if (count($this->entities) == 0) {
            return 1;
        }

        end($this->entities);
        $lastKey = key($this->entities);

        return $lastKey + 1;
    }
}
