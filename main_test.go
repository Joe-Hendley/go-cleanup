package main

import (
	"fmt"
	"testing"
)

func helper(t *testing.T) {
	t.Helper()
	fmt.Println("Helper Called!")
	t.Cleanup(func() {
		fmt.Println("Helper Cleanup Called!")
	})
}

func TestDefer(t *testing.T) {
	fmt.Println("Start Test")
	helper(t)

	fmt.Println("Something that needs deferring called!")
	defer func() { fmt.Println("Defer Called!") }()

	helper(t)
}

func TestCleanup(t *testing.T) {
	fmt.Println("Start Test")
	helper(t)

	fmt.Println("Something that needs cleaning up called!")
	t.Cleanup(func() { fmt.Println("Cleanup Called!") })

	helper(t)
}
