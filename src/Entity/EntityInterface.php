<?php
namespace pdt256\vbscraper\Entity;

interface EntityInterface
{
    /** @return int */
    public function getId();

    /** @var $id */
    public function setId($id);
}
