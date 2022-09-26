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

// === RUN   TestDefer
// Start Test
// Helper Called!
// Something that needs deferring called!
// Helper Called!
// Defer Called!
// Helper Cleanup Called!
// Helper Cleanup Called!

func TestDefer(t *testing.T) {
	fmt.Println("Start Test")
	helper(t)

	fmt.Println("Something that needs deferring called!")
	defer func() { fmt.Println("Defer Called!") }()

	helper(t)
}

// === RUN   TestCleanup
// Start Test
// Helper Called!
// Something that needs cleaning up called!
// Helper Called!
// Helper Cleanup Called!
// Cleanup Called!
// Helper Cleanup Called!

func TestCleanup(t *testing.T) {
	fmt.Println("Start Test")
	helper(t)

	fmt.Println("Something that needs cleaning up called!")
	t.Cleanup(func() { fmt.Println("Cleanup Called!") })

	helper(t)
}
