<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\SetScore;
use pdt256\vbscraper\Entity\Team;
use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\tests\Helper;

class MatchRepositoryTest extends Helper\DoctrineTestCase
{
    protected $metaDataClassNames = [
        'vbscraper:Match',
        'vbscraper:Player',
        'vbscraper:Team',
    ];

    /** @var MatchInterface */
    protected $matchRepository;

    public function setUp()
    {
        $this->matchRepository = $this->repository()->getMatch();
    }

    public function testCRUD()
    {
        $match = $this->setupMatch();

        $this->assertSame(null, $match->getId());
        $this->matchRepository->create($match);
        $this->assertSame(1, $match->getId());

        $match->setGameTimeLengthInSeconds(1);

        $this->assertSame(null, $match->getUpdated());
        $this->matchRepository->update($match);
        $this->assertTrue($match->getUpdated() instanceof \DateTime);

        $this->matchRepository->delete($match);
        $this->assertSame(null, $match->getId());
    }

    public function testFind()
    {
        $this->setupMatch();

        $this->setCountLogger();

        $match = $this->matchRepository->find(1);

        $teamA = $match->getTeamA();
        $teamA->getPlayerA()->getCreated();
        $teamA->getPlayerB()->getCreated();

        $teamB = $match->getTeamB();
        $teamB->getPlayerA()->getCreated();
        $teamB->getPlayerB()->getCreated();

        $this->assertTrue($match instanceof Match);
        $this->assertSame(2, $this->countSQLLogger->getTotalQueries());
    }

    private function setupMatch()
    {
        $player1 = new Player(1, 'John Doe');
        $player2 = new Player(2, 'James Doe');
        $player3 = new Player(3, 'John Smith');
        $player4 = new Player(4, 'James Smith');

        $teamA = new Team;
        $teamA->setPlayerA($player1);
        $teamA->setPlayerB($player2);

        $teamB = new Team;
        $teamB->setPlayerA($player3);
        $teamB->setPlayerB($player4);

        $setScore = new SetScore;
        $setScore->setScoresByString('21-18');

        $match = new Match;
        $match->setTeamA($teamA);
        $match->setTeamB($teamB);
        $match->addSetScore($setScore);

        $this->entityManager->persist($player1);
        $this->entityManager->persist($player2);
        $this->entityManager->persist($player3);
        $this->entityManager->persist($player4);

        $this->entityManager->persist($teamA);
        $this->entityManager->persist($teamB);

        $this->matchRepository->create($match);

        $this->entityManager->flush();
        $this->entityManager->clear();

        return $match;
    }
}
