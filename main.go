package gocleanups

// cleanups struct
type cleanups struct {
	cleanups []func()
}

// Run all cleanups
func (c *cleanups) Run() {
	c.run()
}

// Add cleanup func to cleanups
func (c *cleanups) Add(f func()) {
	c.cleanups = append(c.cleanups, f)
}

// Export cleanups as simple func
func (c *cleanups) Export() func() {
	return c.run
}

func (c *cleanups) run() {
	for _, cleanup := range c.cleanups {
		cleanup()
	}
}

// NewCleanups constructor
func NewCleanups() *cleanups {
	return &cleanups{
		cleanups: make([]func(), 0),
	}
}
