<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\tests\Helper;

class MatchTest extends Helper\DoctrineTestCase
{
    protected $metaDataClassNames = [
        'vbscraper:Match',
    ];

    /** @var MatchInterface */
    protected $matchRepository;

    public function setUp()
    {
        $this->matchRepository = $this->repository()->getMatch();
    }

    public function testCRUD()
    {
        $match = $this->getDummyMatch();

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
}
