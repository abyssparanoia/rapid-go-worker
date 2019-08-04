package main

func main() {
	// Environment
	e := &Environment{}
	e.Get()

	// Dependency
	d := &Dependency{}
	d.Inject(e)

	// Run
	d.Subscription.Listen(d.Handle.Handle)
}
