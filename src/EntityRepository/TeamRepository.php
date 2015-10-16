<?php
namespace pdt256\vbscraper\EntityRepository;

class TeamRepository extends AbstractEntityRepository implements TeamRepositoryInterface
{
    public function findOneById($id)
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

        return $this->returnOrThrowNotFoundException($team);
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

        return $this->returnOrThrowNotFoundException($team);
    }
}
