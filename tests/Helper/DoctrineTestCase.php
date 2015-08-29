<?php
namespace pdt256\vbscraper\tests\Helper;

use Doctrine;
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
        $paths = [__DIR__ . '/../Entity'];
        $isDevMode = true;

        $config = Doctrine\ORM\Tools\Setup::createAnnotationMetadataConfiguration($paths, $isDevMode);
        $xmlDriver = new Doctrine\ORM\Mapping\Driver\XmlDriver(realpath(__DIR__ . '/../../src/Doctrine/Mapping'));
        $config->setMetadataDriverImpl($xmlDriver);
        $config->addEntityNamespace('vbscraper', 'pdt256\vbscraper\Entity');

        $cacheDriver = new Doctrine\Common\Cache\ArrayCache;
        if ($cacheDriver !== null) {
            $config->setMetadataCacheImpl($cacheDriver);
            $config->setQueryCacheImpl($cacheDriver);
            $config->setResultCacheImpl($cacheDriver);
        }

        $dbParams = [
            'driver' => 'pdo_sqlite',
            'memory' => true,
        ];

        $this->entityManager = Doctrine\ORM\EntityManager::create($dbParams, $config);
        $this->entityManagerConfiguration = $this->entityManager->getConnection()->getConfiguration();
        $this->entityManagerConfiguration->setAutoCommit(false);
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
        $this->entityManagerConfiguration->setSQLLogger($sqlLogger);
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
}
