package gocleanups

// Cleanups struct
type Cleanups struct {
	funcs []func()
}

// Run all Cleanups
func (c *Cleanups) Run() {
	c.run()
}

// RunAndReset runs all Cleanups and resets funcs to initial value
func (c *Cleanups) RunAndReset() {
	c.run()
	c.funcs = NewCleanups().funcs
}

// Add cleanup func to Cleanups
func (c *Cleanups) Add(f func()) {
	c.funcs = append(c.funcs, f)
}

// Export Cleanups as simple func
func (c *Cleanups) Export() func() {
	return c.run
}

func (c *Cleanups) run() {
	for _, cleanup := range c.funcs {
		cleanup()
	}
}

// NewCleanups constructor
func NewCleanups() *Cleanups {
	return &Cleanups{
		funcs: make([]func(), 0),
	}
}
