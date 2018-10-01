package testdata

import (
	"github.com/pdt256/vbratings/graphql/generator/testdata/sample"
)

type SimpleDomainQueries struct{}

func (d *SimpleDomainQueries) QueryWithNoParams() bool {
	return true
}

// Line 1
// Line 2
func (d *SimpleDomainQueries) QueryWithDoc() bool {
	return true
}

func (d *SimpleDomainQueries) QueryWithParams(oneInt int, twoString string, threeBool bool) bool {
	return true
}

func (d *SimpleDomainQueries) QueryWithArrayIdentReturn() []bool {
	return []bool{true, false}
}

func (d *SimpleDomainQueries) QueryWithStructReturn() SimpleStruct {
	return SimpleStruct{}
}

func (d *SimpleDomainQueries) QueryWithSelectorStructReturn() sample.SimpleStruct {
	return sample.SimpleStruct{}
}

func (d *SimpleDomainQueries) QueryWithArrayStructReturn() []SimpleStruct {
	return []SimpleStruct{}
}

func (d *SimpleDomainQueries) QueryWithArraySelectorStructReturn() []sample.SimpleStruct {
	return []sample.SimpleStruct{}
}

type SimpleStruct struct {
	Name string
}
