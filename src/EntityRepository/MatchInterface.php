<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity;

interface MatchInterface
{
    public function create(Entity\Match & $attribute);
    public function update(Entity\Match & $attribute);
    public function delete(Entity\Match & $attribute);

    /**
     * @param int $id
     * @return Entity\Match
     */
    public function find($id);
}
