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

$matchImporter = new MatchImporter(
    $doctrine->repository()->getMatch(),
    $doctrine->repository()->getTeam(),
    $doctrine->repository()->getPlayer()
);

$bvbBase = 'http://bvbinfo.com/';

$content = file_get_contents($bvbBase . 'season.asp');
$seasonUrls = BvbInfoScraper::getSeasonUrls($content);

foreach (array_reverse($seasonUrls) as $seasonUrl) {
    $seasonUrl = $bvbBase . $seasonUrl;

    $year = (int) substr($seasonUrl, -4);
    if ($year < 2015) {
        continue;
    }

    echo 'Season: ' . $seasonUrl . PHP_EOL;

    $content = file_get_contents($seasonUrl);
    $seasonTournamentUrls = BvbInfoScraper::getSeasonTournamentUrls($content);

//    $seasonUrl = $bvbBase . '/Season.asp?AssocID=1&Year=2015';
//    $content = file_get_contents($seasonUrl);
//    $seasonTournamentUrls = BvbInfoScraper::getSeasonTournamentUrls($content);

    foreach ($seasonTournamentUrls as $seasonTournamentUrl) {
        $seasonTournamentUrl = $bvbBase . $seasonTournamentUrl . '&Process=Matches';
        echo 'Tournament: ' . $seasonTournamentUrl . PHP_EOL;

        $content = file_get_contents($seasonTournamentUrl);
        $matches = BvbInfoScraper::getMatches($content);

        $matchImporter->import($matches);
    }
}
