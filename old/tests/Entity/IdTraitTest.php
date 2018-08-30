<?php
namespace pdt256\vbscraper\Entity;

class IdTraitTest extends \PHPUnit_Framework_TestCase
{
    public function testId()
    {
        $mock = $this->getObjectForTrait('pdt256\vbscraper\Entity\IdTrait');
        $mock->setId(1);
        $this->assertSame(1, $mock->getId());
    }
}
