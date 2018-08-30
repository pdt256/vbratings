<?php
namespace pdt256\vbscraper\EntityRepository;

use pdt256\vbscraper\Entity\Player;

/**
 * @method Player findOneById($id)
 */
interface PlayerRepositoryInterface extends AbstractRepositoryInterface
{
    /**
     * @param int $vbId
     * @return Player
     * @throws EntityNotFoundException
     */
    public function findOneByVbId($vbId);
}
