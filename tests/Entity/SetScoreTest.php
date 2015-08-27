<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\Exception\ValidatorException;
use Symfony\Component\Validator\Validation;

class SetScoreTest extends \PHPUnit_Framework_TestCase
{
    public function testFailsValidationWithWrongScore()
    {
        $setScore = new SetScore;
        $setScore->setTeamAScore(-1);
        $setScore->setTeamBScore(-1);

        $validator = Validation::createValidatorBuilder()
            ->addMethodMapping('loadValidatorMetadata')
            ->getValidator();

        $validationErrors = $validator->validate($setScore);

        $this->assertSame(2, $validationErrors->count());
        $this->assertSame('teamAScore', $validationErrors->get(0)->getPropertyPath());
        $this->assertSame('teamBScore', $validationErrors->get(1)->getPropertyPath());
    }

    public function testPassesValidationWithNullScoreWithForfeit()
    {
        $setScore = new SetScore;
        $setScore->setTeamBForfeit();

        $validator = Validation::createValidatorBuilder()
            ->addMethodMapping('loadValidatorMetadata')
            ->getValidator();

        $this->assertEmpty($validator->validate($setScore));
    }
}
