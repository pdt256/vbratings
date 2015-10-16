<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\tests\Helper;

class PlayerRepositoryTest extends Helper\DoctrineTestCase
{
    protected $metaDataClassNames = [
        'vbscraper:Player',
        'vbscraper:Team',
    ];

    /** @var PlayerRepositoryInterface */
    protected $playerRepository;

    public function setUp()
    {
        $this->playerRepository = $this->repository()->getPlayer();
    }

    public function testCRUD()
    {
        $player = $this->setupPlayer();

        $this->assertSame(1, $player->getId());

        $player->setName('James Doe');

        $this->assertSame(null, $player->getUpdated());
        $this->playerRepository->update($player);
        $this->assertTrue($player->getUpdated() instanceof \DateTime);

        $this->playerRepository->delete($player);
        $this->assertSame(null, $player->getId());
    }

    public function testFind()
    {
        $this->setupPlayer();
        $this->entityManager->clear();

        $this->setCountLogger();

        $player = $this->playerRepository->findOneById(1);

        $this->assertTrue($player instanceof Player);
        $this->assertSame(1, $this->getTotalQueries());
    }

    private function setupPlayer()
    {
        $player = new Player(123, 'John Doe');

        $this->playerRepository->create($player);

        return $player;
    }
}
