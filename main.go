package main

import "fmt"

type BrowserNavigator struct {
	current       string
	backStack     []string
	forwardStack  []string
}

func Constructor(startPage string) BrowserNavigator {
	return BrowserNavigator{
		current:      startPage,
		backStack:    []string{},
		forwardStack: []string{},
	}
}

func (b *BrowserNavigator) hopTo(page string) {
	b.backStack = append(b.backStack, b.current)
	b.current = page
	b.forwardStack = []string{} // clear forward history
}

func (b *BrowserNavigator) backtrack(steps int) string {
	for steps > 0 && len(b.backStack) > 0 {
		b.forwardStack = append(b.forwardStack, b.current)
		n := len(b.backStack)
		b.current = b.backStack[n-1]
		b.backStack = b.backStack[:n-1]
		steps--
	}
	return b.current
}

func (b *BrowserNavigator) leapForward(steps int) string {
	for steps > 0 && len(b.forwardStack) > 0 {
		b.backStack = append(b.backStack, b.current)
		n := len(b.forwardStack)
		b.current = b.forwardStack[n-1]
		b.forwardStack = b.forwardStack[:n-1]
		steps--
	}
	return b.current
}
package main

import "fmt"

type BrowserNavigator struct {
	current       string
	backStack     []string
	forwardStack  []string
}

func Constructor(startPage string) BrowserNavigator {
	return BrowserNavigator{
		current:      startPage,
		backStack:    []string{},
		forwardStack: []string{},
	}
}

func (b *BrowserNavigator) hopTo(page string) {
	b.backStack = append(b.backStack, b.current)
	b.current = page
	b.forwardStack = []string{} // clear forward history
}

func (b *BrowserNavigator) backtrack(steps int) string {
	for steps > 0 && len(b.backStack) > 0 {
		b.forwardStack = append(b.forwardStack, b.current)
		n := len(b.backStack)
		b.current = b.backStack[n-1]
		b.backStack = b.backStack[:n-1]
		steps--
	}
	return b.current
}

func (b *BrowserNavigator) leapForward(steps int) string {
	for steps > 0 && len(b.forwardStack) > 0 {
		b.backStack = append(b.backStack, b.current)
		n := len(b.forwardStack)
		b.current = b.forwardStack[n-1]
		b.forwardStack = b.forwardStack[:n-1]
		steps--
	}
	return b.current
}

func main() {
	b := Constructor("home")
	fmt.Println("Starting page:", b.current)

	b.hopTo("p1")
	fmt.Println("Current page:", b.current)

	b.hopTo("p2")
	fmt.Println("Current page:", b.current)

	b.hopTo("p3")
	fmt.Println("Current page:", b.current)

	fmt.Println("Backtrack(5):", b.backtrack(5)) // <- solo puede ir 3 atrás, debe quedar en "home"
	fmt.Println("LeapForward(10):", b.leapForward(10)) // <- debe llegar hasta "p3"

	b.hopTo("new") // esto borra el forward stack
	fmt.Println("Current page after hopTo('new'):", b.current)

	fmt.Println("LeapForward(1) (should not change):", b.leapForward(1)) // <- no debe avanzar
}


