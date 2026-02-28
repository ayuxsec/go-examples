package main

import "fmt"

// A "Greeter" that knows how to greet
type Greeter struct {
	Prefix string
}

// Greeter has a method called SayHello
func (g *Greeter) SayHello(name string) string {
	return g.Prefix + ", " + name + "!"
}

// Another struct that accepts a function as a field
type EventHandler struct {
	OnUserJoin func(name string) string
}

func main() {
	// Using the anonymous struct pattern (like http.Transport)
	handler1 := &EventHandler{
		OnUserJoin: (&Greeter{
			Prefix: "Welcome",
		}).SayHello,
	}

	// Identical but more explicit
	greeter := &Greeter{Prefix: "Welcome"}
	handler2 := &EventHandler{
		OnUserJoin: greeter.SayHello,
	}

	fmt.Println(handler1.OnUserJoin("Alice")) // → Welcome, Alice!
	fmt.Println(handler2.OnUserJoin("bob"))   // → Welcome, bob!
}
