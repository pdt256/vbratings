<?php
namespace pdt256\vbscraper\Entity;

use Symfony\Component\Validator\Mapping\ClassMetadata;

interface ValidationInterface
{
    public static function loadValidatorMetadata(ClassMetadata $metadata);
}
