<?php
use Doctrine\Common\Cache\ArrayCache;
use pdt256\vbscraper\Lib\DoctrineHelper;
use pdt256\vbscraper\Service\BvbInfoScraper;
use pdt256\vbscraper\Service\Import\MatchImporter;

error_reporting(E_ALL);
date_default_timezone_set('UTC');
ini_set('memory_limit', '256M');

require_once __DIR__ . '/../vendor/autoload.php';

$doctrine = new DoctrineHelper(new ArrayCache);
$doctrine->setup([
    'driver' => 'pdo_sqlite',
    'path' => __DIR__ . '/../data/db.sqlite',
]);

$doctrine->clearDB();

class LiveImport
{
    private $bvbBase = 'http://bvbinfo.com/';

    /** @var MatchImporter */
    private $matchImporter;

    public function __construct(DoctrineHelper $doctrine)
    {
        $this->matchImporter = new MatchImporter(
            $doctrine->repository()->getMatch(),
            $doctrine->repository()->getTeam(),
            $doctrine->repository()->getPlayer()
        );
    }

    public function execute()
    {
        $content = file_get_contents($this->bvbBase . 'season.asp');
        $seasonUrls = BvbInfoScraper::getSeasonUrls($content);

        foreach (array_reverse($seasonUrls) as $seasonUrl) {
            $seasonUrl = $this->bvbBase . $seasonUrl;
            $this->importSeason($seasonUrl);
        }
    }

    /**
     * @param $seasonTournamentUrls
     */
    public function importSeasonTournamentUrls($seasonTournamentUrls)
    {
        foreach ($seasonTournamentUrls as $seasonTournamentUrl) {
            $seasonTournamentUrl = $this->bvbBase . $seasonTournamentUrl . '&Process=Matches';
            $this->importTournamentMatches($seasonTournamentUrl);
        }
    }

    /**
     * @param $seasonTournamentUrl
     */
    public function importTournamentMatches($seasonTournamentUrl)
    {
        echo 'Tournament: ' . $seasonTournamentUrl . PHP_EOL;

        $content = file_get_contents($seasonTournamentUrl);
        $matches = BvbInfoScraper::getMatches($content);

        $this->matchImporter->import($matches);
    }

    private function yearIsValidInSeasonUrl($seasonUrl)
    {
        $year = (int) substr($seasonUrl, -4);
        return ($year >= 2015);
    }

    public function importAVP2015()
    {
        $seasonUrl = $this->bvbBase . '/Season.asp?AssocID=1&Year=2015';
        $this->importSeason($seasonUrl);
    }

    /**
     * @param $seasonUrl
     */
    private function importSeason($seasonUrl)
    {
        if (! $this->yearIsValidInSeasonUrl($seasonUrl)) {
            return;
        }

        echo 'Season: ' . $seasonUrl . PHP_EOL;

        $content = file_get_contents($seasonUrl);
        $seasonTournamentUrls = BvbInfoScraper::getSeasonTournamentUrls($content);

        $this->importSeasonTournamentUrls($seasonTournamentUrls);
    }
}

$liveImport = new LiveImport($doctrine);
//$liveImport->execute();
$liveImport->importAVP2015();
