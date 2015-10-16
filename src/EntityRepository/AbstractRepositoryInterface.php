<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\EntityInterface;

interface AbstractRepositoryInterface
{
    public function getQueryBuilder();
    public function create(EntityInterface & $entity);
    public function update(EntityInterface & $entity);
    public function delete(EntityInterface $entity);
    public function persist(EntityInterface & $entity);
    public function flush();

    /**
     * @param int $id
     * @return EntityInterface
     * @throws EntityNotFoundException
     */
    public function findOneById($id);
}
