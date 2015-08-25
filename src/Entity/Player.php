<?php
namespace pdt256\vbscraper\Entity;

class Player
{
    /** @var string */
    private $name;

    /** @var int */
    private $vbId;

    /**
     * @param int $vbId
     * @param string $name
     */
    public function __construct($vbId, $name)
    {
        $this->setVbId($vbId);
        $this->setName($name);
    }

    public function getName()
    {
        return $this->name;
    }

    /**
     * @param string $name
     */
    public function setName($name)
    {
        $this->name = (string) $name;
    }

    public function getVbId()
    {
        return $this->vbId;
    }

    /**
     * @param int $vbId
     */
    public function setVbId($vbId)
    {
        $this->vbId = (int) $vbId;
    }
}
