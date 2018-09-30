package testdata

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

func (d *SimpleDomainQueries) QueryWithArrayReturn() []bool {
	return []bool{true, false}
}
