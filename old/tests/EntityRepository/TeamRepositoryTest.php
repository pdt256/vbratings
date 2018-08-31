<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\Team;
use pdt256\vbscraper\tests\Helper;

class TeamRepositoryTest extends Helper\DoctrineTestCase
{
    protected $metaDataClassNames = [
        'vbscraper:Team',
        'vbscraper:Player',
    ];

    /** @var TeamRepositoryInterface */
    protected $teamRepository;

    public function setUp()
    {
        $this->teamRepository = $this->repository()->getTeam();
    }

    public function testCRUD()
    {
        $team = $this->setupTeam();

        $this->assertSame(1, $team->getId());

        $playerA = new Player(123, 'John Doe');
        $this->entityManager->persist($playerA);
        $this->entityManager->flush();

        $team->setPlayerA($playerA);

        $this->assertSame(null, $team->getUpdated());
        $this->teamRepository->update($team);
        $this->assertTrue($team->getUpdated() instanceof \DateTime);

        $this->teamRepository->delete($team);
        $this->assertSame(null, $team->getId());
    }

    public function testFindOneById()
    {
        $this->setupTeam();
        $this->entityManager->clear();

        $this->setCountLogger();

        $team = $this->teamRepository->findOneById(1);

        $this->assertTrue($team instanceof Team);
        $this->assertSame(1, $this->getTotalQueries());
    }

    /**
     * @expectedException \pdt256\vbscraper\EntityRepository\EntityNotFoundException
     * @expectedExceptionMessage Team not found
     */
    public function testFindOneByIdThrowsException()
    {
        $team = $this->teamRepository->findOneById(1);
    }

    private function setupTeam()
    {
        $team = new Team;
        $this->teamRepository->create($team);
        return $team;
    }
}