package testdata

type SimpleDomainCommands struct{}

func (d *SimpleDomainCommands) CommandWithNoReturn() {}

func (d *SimpleDomainCommands) CommandReturnsError() error {
	return nil
}

func (d *SimpleDomainCommands) CommandWithParams(oneInt int, twoString string, threeBool bool) {}
