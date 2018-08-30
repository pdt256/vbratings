<?php
namespace pdt256\vbscraper\Service\Import;

use pdt256\vbscraper\Entity\Player;
use pdt256\vbscraper\Entity\Match;
use pdt256\vbscraper\Entity\Team;
use pdt256\vbscraper\EntityRepository\EntityNotFoundException;
use pdt256\vbscraper\EntityRepository\MatchRepositoryInterface;
use pdt256\vbscraper\EntityRepository\PlayerRepositoryInterface;
use pdt256\vbscraper\EntityRepository\TeamRepositoryInterface;
use Symfony\Component\Validator\Exception\ValidatorException;
use Symfony\Component\Validator\Validation;

class MatchImporter
{
    /** @var MatchRepositoryInterface */
    private $matchRepository;

    /** @var TeamRepositoryInterface */
    private $teamRepository;

    /** @var PlayerRepositoryInterface */
    private $playerRepository;

    public function __construct(
        MatchRepositoryInterface $matchRepository,
        TeamRepositoryInterface $teamRepository,
        PlayerRepositoryInterface $playerRepository
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

        try {
            return $this->teamRepository->findOneByPlayers($team->getPlayerA(), $team->getPlayerB());
        } catch (EntityNotFoundException $e) {
            $this->teamRepository->persist($team);
            return $team;
        }
    }

    /**
     * @param Player $player
     * @return Player
     */
    private function createPlayerIfNotFound($player)
    {
        try {
            return $this->playerRepository->findOneByVbId($player->getVbId());
        } catch (EntityNotFoundException $e) {
            $this->playerRepository->persist($player);
            return $player;
        }
    }
}
