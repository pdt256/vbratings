<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Team;

class TeamRepository extends AbstractEntityRepository implements TeamInterface
{
    public function create(Team & $entity)
    {
        $this->createEntity($entity);
    }

    public function update(Team & $entity)
    {
        $this->updateEntity($entity);
    }

    public function delete(Team & $entity)
    {
        $this->deleteEntity($entity);
    }

    public function find($id)
    {
        $qb = $this->getQueryBuilder();

        $team = $qb
            ->select('team')
            ->from('vbscraper:Team', 'team')

            ->leftJoin('team.playerA', 'playerA')
            ->leftJoin('team.playerB', 'playerB')

            ->where('team.id = :id')->setParameter('id', $id)
            ->getQuery()
            ->getOneOrNullResult();

        return $team;
    }

    public function findOneByPlayers($playerA, $playerB)
    {
        $qb = $this->getQueryBuilder();

        $team = $qb
            ->select('team')
            ->from('vbscraper:Team', 'team')

            ->leftJoin('team.playerA', 'playerA')
            ->leftJoin('team.playerB', 'playerB')

            ->where('playerA.id = :playerAId')->setParameter('playerAId', $playerA->getId())
            ->andWhere('playerB.id = :playerBId')->setParameter('playerBId', $playerB->getId())
            ->getQuery()
            ->getOneOrNullResult();

        return $team;
    }
}
