package testdata

type SingleDomainSingleQuery struct{}

// Line 1
// Line 2
func (d *SingleDomainSingleQuery) GetQuery(oneInt int, twoString string, threeBool bool) bool {
	return true
}
