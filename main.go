package gocleanups

// Cleanups struct
type Cleanups []func()

// Len returns the number of cleanup functions
func (c Cleanups) Len() int {
	return len(c)
}

// Run all Cleanups
func (c Cleanups) Run() {
	for _, f := range c {
		f()
	}
}

// RunAndReset runs all Cleanups and resets funcs to empty slice
func (c *Cleanups) RunAndReset() {
	c.Run()
	*c = NewCleanups()
}

// Add cleanup funcs to Cleanups
func (c *Cleanups) Add(funcs ...func()) {
	*c = append(*c, funcs...)
}

// Export Cleanups as simple func
func (c Cleanups) Export() func() {
	return c.Run
}

// NewCleanups constructor
func NewCleanups(funcs ...func()) Cleanups {
	return append(make([]func(), 0, len(funcs)), funcs...)
}
