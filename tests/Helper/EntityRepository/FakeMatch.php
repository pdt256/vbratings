<?php
namespace pdt256\vbscraper\tests\Helper\EntityRepository;

use pdt256\vbscraper\EntityRepository\MatchRepositoryInterface;
use pdt256\vbscraper\Entity\Match;

class FakeMatch extends AbstractFake implements MatchRepositoryInterface
{
    public function __construct()
    {
        $this->setReturnValue(new Match);
    }

}
