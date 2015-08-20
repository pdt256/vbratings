Volleyball Score Scraper
========================

## Introduction


## Description


## Installation

Add the following lines to your ``composer.json`` file.

```JSON
{
    "require": {
        "pdt256/vbscraper": "dev-master"
    }
}
```

```
   composer install
```

## Unit Tests:

<pre>
    vendor/bin/phpunit
</pre>

### With Code Coverage:

<pre>
    vendor/bin/phpunit --coverage-text --coverage-html coverage_report
</pre>

## Run Coding Standards Test:

<pre>
    vendor/bin/phpcs --standard=PSR2 src/ tests/
</pre>


## License

```
Copyright 2015 Jamie Isaacs pdt256@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
