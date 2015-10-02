<?php
namespace pdt256\vbscraper\Lib;

use Doctrine;
use Doctrine\ORM\Tools\SchemaTool;

class DoctrineHelper
{
    /** @var Doctrine\ORM\EntityManager */
    protected $entityManager;

    /** @var Doctrine\DBAL\Configuration */
    protected $entityManagerConfiguration;

    /** @var Doctrine\Common\Cache\CacheProvider */
    protected $cacheDriver;

    /** @var Doctrine\DBAL\Configuration */
    protected $config;

    /** @var SchemaTool */
    protected $schemaTool;

    public function __construct(Doctrine\Common\Cache\CacheProvider $cacheDriver = null)
    {
        $paths = [__DIR__ . '/../Entity'];
        $isDevMode = true;

        $this->config = Doctrine\ORM\Tools\Setup::createAnnotationMetadataConfiguration($paths, $isDevMode);
        $xmlDriver = new Doctrine\ORM\Mapping\Driver\XmlDriver(realpath(__DIR__ . '/../Doctrine/Mapping'));
        $this->config->setMetadataDriverImpl($xmlDriver);
        $this->config->addEntityNamespace('vbscraper', 'pdt256\vbscraper\Entity');

        if ($cacheDriver !== null) {
            $this->cacheDriver = $cacheDriver;
            $this->config->setMetadataCacheImpl($this->cacheDriver);
            $this->config->setQueryCacheImpl($this->cacheDriver);
            $this->config->setResultCacheImpl($this->cacheDriver);
        }
    }

    public function clearCache()
    {
        $this->cacheDriver->deleteAll();
    }

    public function getCacheDriver()
    {
        return $this->cacheDriver;
    }

    public function getEntityManager()
    {
        return $this->entityManager;
    }

    public function setSqlLogger(Doctrine\DBAL\Logging\SQLLogger $sqlLogger)
    {
        $this->entityManagerConfiguration->setSQLLogger($sqlLogger);
    }

    public function setup(array $dbParams)
    {
        $this->entityManager = Doctrine\ORM\EntityManager::create($dbParams, $this->config);
        $this->entityManagerConfiguration = $this->entityManager->getConnection()->getConfiguration();
        $this->schemaTool = new Doctrine\ORM\Tools\SchemaTool($this->entityManager);
    }

    public function initializeSchema()
    {
        $classes = $this->entityManager->getMetaDataFactory()->getAllMetaData();
        $this->schemaTool->updateSchema($classes);
    }

    public function repository()
    {
        return new FactoryRepository($this->entityManager);
    }

    public function dropSchema()
    {
        $classes = $this->entityManager->getMetaDataFactory()->getAllMetaData();
        $this->schemaTool->dropSchema($classes);
    }

    public function clearDB()
    {
        $this->initializeSchema();
        $this->dropSchema();
        $this->initializeSchema();
    }
}
