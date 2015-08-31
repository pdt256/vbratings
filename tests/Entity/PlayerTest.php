<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\Validation;

class PlayerTest extends \PHPUnit_Framework_TestCase
{
    public function testCreate()
    {
        $player = new Player(123, 'John Doe');

        $validator = Validation::createValidatorBuilder()
            ->addMethodMapping('loadValidatorMetadata')
            ->getValidator();

        $this->assertEmpty($validator->validate($player));
        $this->assertSame('John Doe', $player->getname());
        $this->assertSame(123, $player->getVbId());
    }
}
