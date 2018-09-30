package testdata

type SimpleDomain struct{}

// Query 1 Doc
func (d *SimpleDomain) Query1() bool { return true }

// Query 2 Doc
// Second Line
func (d *SimpleDomain) Query2() (bool, error) { return true, nil }

// Command 1 Doc
func (d *SimpleDomain) Command1() {}

// Command 2 Doc
func (d *SimpleDomain) Command2() error { return nil }
