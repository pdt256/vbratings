<?php
namespace pdt256\vbscraper\Entity;

class CreatedTraitTest extends \PHPUnit_Framework_TestCase
{
    public function testCreate()
    {
        $mock = $this->getObjectForTrait('pdt256\vbscraper\Entity\CreatedTrait');
        $mock->setCreated(new \DateTime);
        $this->assertTrue($mock->getCreated() instanceof \DateTime);
    }

    public function testSetCreatedWithNull()
    {
        $mock = $this->getObjectForTrait('pdt256\vbscraper\Entity\CreatedTrait');
        $mock->setCreated();
        $this->assertTrue($mock->getCreated() instanceof \DateTime);
    }
}
