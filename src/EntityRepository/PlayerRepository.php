<?php
namespace pdt256\vbscraper\EntityRepository;

class PlayerRepository extends AbstractEntityRepository implements PlayerRepositoryInterface
{
    public function findOneByVbId($vbId)
    {
        return $this->returnOrThrowNotFoundException(
            $this->findOneBy(['vbId' => $vbId])
        );
    }
}
