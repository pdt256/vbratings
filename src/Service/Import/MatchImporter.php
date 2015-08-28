<?php
namespace pdt256\vbscraper\Service\Import;

use pdt256\vbscraper\Entity;
use Symfony\Component\Validator\Exception\ValidatorException;
use Symfony\Component\Validator\Validation;

class MatchImporter
{
    /**
     * @param Entity\Match[] $matches
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

                // $this->matchRepository->create($match);
                $importResult->incrementSuccess();
            } catch (\Exception $e) {
                $importResult->addFailedRow($match);
                $importResult->addErrorMessage($e->getMessage());
            }
        }

        return $importResult;
    }
}