package bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *ByteCounter) Read(p []byte) (n int, err error) {
	// TODO implement me
	panic("implement me")
}
