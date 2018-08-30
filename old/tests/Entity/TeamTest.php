<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\Validation;

class TeamTest extends \PHPUnit_Framework_TestCase
{
    public function testCreate()
    {
        $team = new Team;
        $team->setPlayerA(new Player(1, 'John Doe'));
        $team->setPlayerB(new Player(2, 'James Doe'));

        $validator = Validation::createValidatorBuilder()
            ->addMethodMapping('loadValidatorMetadata')
            ->getValidator();

        $this->assertEmpty($validator->validate($team));
        $this->assertTrue($team->getPlayerA() instanceof Player);
        $this->assertTrue($team->getPlayerB() instanceof Player);
        $this->assertSame('[(1) John Doe, (2) James Doe]', $team->__toString());
    }
}
