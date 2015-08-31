<?php
namespace pdt256\vbscraper\Entity;

class UpdatedTest extends \PHPUnit_Framework_TestCase
{
    public function testCreate()
    {
        $mock = $this->getObjectForTrait('pdt256\vbscraper\Entity\UpdatedTrait');
        $mock->setUpdated(new \DateTime);
        $mock->preUpdate();
        $this->assertTrue($mock->getUpdated() instanceof \DateTime);
    }

    public function testSetUpdatedWithNull()
    {
        $mock = $this->getObjectForTrait('pdt256\vbscraper\Entity\UpdatedTrait');
        $mock->setUpdated();
        $this->assertTrue($mock->getUpdated() instanceof \DateTime);
    }

    public function testGetUpdatedWithNull()
    {
        $mock = $this->getObjectForTrait('pdt256\vbscraper\Entity\UpdatedTrait');
        $this->assertSame(null, $mock->getUpdated());
    }
}
