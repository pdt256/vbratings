<?php
namespace pdt256\vbscraper\tests\Helper\EntityRepository;

use pdt256\vbscraper\Entity;
use Exception;

class AbstractFake
{
    /** @var Entity\EntityInterface */
    public $returnValue;

    /** @var Exception|null */
    public $crudExceptionToThrow;

    protected function getReturnValue()
    {
        return $this->returnValue;
    }

    protected function getReturnValueAsArray()
    {
        if ($this->returnValue === null) {
            return [];
        }

        return [$this->returnValue];
    }

    public function setReturnValue(Entity\EntityInterface $returnValue = null)
    {
        $this->returnValue = $returnValue;
    }

    public function setCrudException(Exception $exception)
    {
        $this->crudExceptionToThrow = $exception;
    }

    public function throwCrudExceptionIfSet()
    {
        if ($this->crudExceptionToThrow !== null) {
            throw $this->crudExceptionToThrow;
        }
    }

    public function flush()
    {
    }
}
