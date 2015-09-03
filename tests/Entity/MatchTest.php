<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\Validation;

class MatchTest extends \PHPUnit_Framework_TestCase
{
    public function testCreate()
    {
        $teamA = new Team();
        $teamA->setPlayerA(new Player(1, 'John Doe'));
        $teamA->setPlayerB(new Player(2, 'James Doe'));

        $teamB = new Team();
        $teamB->setPlayerA(new Player(3, 'John Smith'));
        $teamB->setPlayerB(new Player(4, 'James Smith'));

        $setScore = new SetScore;
        $setScore->setScoresByString('21-18');

        $match = new Match;
        $match->setTeamA($teamA);
        $match->setTeamB($teamB);
        $match->addSetScore($setScore);
        $match->addSetScoreByString('21-16');
        $match->setGameTimeLengthInSeconds(5);

        $validator = Validation::createValidatorBuilder()
            ->addMethodMapping('loadValidatorMetadata')
            ->getValidator();

        $this->assertEmpty($validator->validate($match));
        $this->assertTrue($match->getTeamA() instanceof Team);
        $this->assertTrue($match->getTeamB() instanceof Team);
        $this->assertSame(2, count($match->getSetScores()));
        $this->assertTrue($match->getSetscores()[0] instanceof SetSCore);
        $this->assertTrue($match->getSetscore(1) instanceof SetSCore);
        $this->assertSame(null, $match->getSetScore(3));
        $this->assertSame(5, $match->getGameTimeLengthInSeconds());
        $this->assertSame('[(1) John Doe, (2) James Doe] def [(3) John Smith, (4) James Smith]', $match->__toString());
    }

    /**
     * @expectedException \RuntimeException
     * @expectedExceptionMessage Exceeded maximum set scores.
     */
    public function testAddSetScoreThrowsException()
    {
        $match = new Match;
        $match->addSetScore(new SetScore);
        $match->addSetScore(new SetScore);
        $match->addSetScore(new SetScore);
        $match->addSetScore(new SetScore);
    }
}
