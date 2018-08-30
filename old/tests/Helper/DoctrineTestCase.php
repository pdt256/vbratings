<?php
namespace pdt256\vbscraper\tests\Helper;

use Doctrine;
use Doctrine\Common\Cache\ArrayCache;
use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\SetScore;
use pdt256\vbscraper\Entity\Team;
use pdt256\vbscraper\Lib\DoctrineHelper;
use pdt256\vbscraper\Lib\FactoryRepository;

abstract class DoctrineTestCase extends \PHPUnit_Framework_TestCase
{
    /** @var Doctrine\ORM\EntityManager */
    protected $entityManager;

    /** @var Doctrine\DBAL\Configuration */
    protected $entityManagerConfiguration;

    /** @var CountSQLLogger */
    protected $countSQLLogger;

    /** @var string[] */
    protected $metaDataClassNames;

    /** @var DoctrineHelper */
    protected $doctrineHelper;

    public function __construct($name = null, array $data = array(), $dataName = '')
    {
        parent::__construct($name, $data, $dataName);

        if ($this->metaDataClassNames !== null) {
            $this->setupEntityManager();
        }
    }

    protected function setupEntityManager()
    {
        $this->getConnection();
        $this->setupTestSchema();
    }

    private function getConnection()
    {
        $this->doctrineHelper = new DoctrineHelper(new ArrayCache);
        $this->doctrineHelper->setup([
            'driver' => 'pdo_sqlite',
            'memory' => true,
        ]);

        $this->entityManager = $this->doctrineHelper->getEntityManager();
    }

    private function setupTestSchema()
    {
        $this->entityManager->clear();

        if (empty($this->metaDataClassNames)) {
            $classes = $this->entityManager->getMetaDataFactory()->getAllMetaData();
        } else {
            $classes = [];
            foreach ($this->metaDataClassNames as $className) {
                $classes[] = $this->entityManager->getMetaDataFactory()->getMetadataFor($className);
            }
        }

        $tool = new Doctrine\ORM\Tools\SchemaTool($this->entityManager);
        // $tool->dropSchema($classes);
        $tool->createSchema($classes);
    }

    public function setEchoLogger()
    {
        $this->setSqlLogger(new Doctrine\DBAL\Logging\EchoSQLLogger);
    }

    public function setCountLogger()
    {
        $this->countSQLLogger = new CountSQLLogger;
        $this->setSqlLogger($this->countSQLLogger);
    }

    private function setSqlLogger(Doctrine\DBAL\Logging\SQLLogger $sqlLogger)
    {
        $this->doctrineHelper->setSqlLogger($sqlLogger);
    }

    protected function getTotalQueries()
    {
        return $this->countSQLLogger->getTotalQueries();
    }

    protected function repository()
    {
        return new FactoryRepository($this->entityManager);
    }

    protected function beginTransaction()
    {
        $this->entityManager->getConnection()->beginTransaction();
    }

    protected function rollback()
    {
        $this->entityManager->getConnection()->rollback();
    }

    protected function getDummyMatch()
    {
        $teamA = new Team;
        $teamA->setPlayerA(new Player(1, 'John Doe'));
        $teamA->setPlayerB(new Player(2, 'James Doe'));

        $teamB = new Team;
        $teamB->setPlayerA(new Player(3, 'John Smith'));
        $teamB->setPlayerB(new Player(4, 'James Smith'));

        $setScore = new SetScore;
        $setScore->setScoresByString('21-18');

        $match = new Match;
        $match->setTeamA($teamA);
        $match->setTeamB($teamB);
        $match->addSetScore($setScore);

        return $match;
    }
}
