<?php
namespace pdt256\vbscraper\Service\Import;

use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\Entity\Team;
use pdt256\vbscraper\EntityRepository;
use Symfony\Component\Validator\Exception\ValidatorException;
use Symfony\Component\Validator\Validation;

class MatchImporter
{
    /** @var EntityRepository\MatchInterface */
    private $matchRepository;

    /** @var EntityRepository\TeamInterface */
    private $teamRepository;

    /** @var EntityRepository\PlayerInterface */
    private $playerRepository;

    public function __construct(
        EntityRepository\MatchInterface $matchRepository,
        EntityRepository\TeamInterface $teamRepository,
        EntityRepository\PlayerInterface $playerRepository
    ) {
        $this->matchRepository = $matchRepository;
        $this->teamRepository = $teamRepository;
        $this->playerRepository = $playerRepository;
    }

    /**
     * @param Match[] $matches
     * @return MatchImportResult
     */
    public function import(array $matches)
    {
        $validator = Validation::createValidatorBuilder()
            ->addMethodMapping('loadValidatorMetadata')
            ->getValidator();

        $importResult = new MatchImportResult;

        foreach ($matches as $match) {
            try {
                $errors = $validator->validate($match);
                if ($errors->count() > 0) {
                    throw new ValidatorException('Invalid Match ' . (string) $errors);
                }

                $match->setTeamA($this->createTeamIfNotFound($match->getTeamA()));
                $match->setTeamB($this->createTeamIfNotFound($match->getTeamB()));

                $this->matchRepository->persist($match);
                $importResult->incrementSuccess();
            } catch (\Exception $e) {
                $importResult->addFailedRow($match);
                $importResult->addErrorMessage($e->getMessage());
            }
            $this->matchRepository->flush();
        }

        $this->matchRepository->clear();

        return $importResult;
    }

    /**
     * @param Team $team
     * @return Team
     */
    private function createTeamIfNotFound($team)
    {
        $team->setPlayerA($this->createPlayerIfNotFound($team->getPlayerA()));
        $team->setPlayerB($this->createPlayerIfNotFound($team->getPlayerB()));

        $teamEntity = $this->teamRepository->findOneByPlayers($team->getPlayerA(), $team->getPlayerB());

        if ($teamEntity === null) {
            $this->teamRepository->persist($team);
            return $team;
        }

        return $teamEntity;
    }

    /**
     * @param Player $player
     * @return Player
     */
    private function createPlayerIfNotFound($player)
    {
        $playerEntity = $this->playerRepository->findOneByVbId($player->getVbId());

        if ($playerEntity === null) {
            $this->playerRepository->persist($player);
            return $player;
        }

        return $playerEntity;
    }
}
