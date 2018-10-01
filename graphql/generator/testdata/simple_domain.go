package testdata

type SimpleDomain struct{}

// Query 1 Doc
func (d *SimpleDomain) Query1() bool { return true }

// Query 2 Doc
// Second Line
func (d *SimpleDomain) Query2() (string, error) { return "", nil }

func (d *SimpleDomain) Query3(one int, two string, three bool) int { return 1 }

func (d *SimpleDomain) Query4() Struct1 { return Struct1{} }

// Command 1 Doc
func (d *SimpleDomain) Command1() {}

// Command 2 Doc
func (d *SimpleDomain) Command2() error { return nil }

func (d *SimpleDomain) Command3(one int, two string, three bool) {}
