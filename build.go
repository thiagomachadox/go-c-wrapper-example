//go:build ignore

package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Compile C code
	cmd := exec.Command("gcc", "-c", "-fPIC", "math.c", "-o", "math.o")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to compile C code: %v", err)
	}

	// Create shared library
	cmd = exec.Command("gcc", "-shared", "-o", "libmath.so", "math.o")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to create shared library: %v", err)
	}

	// Clean up object file
	if err := os.Remove("math.o"); err != nil {
		log.Printf("Failed to remove math.o: %v", err)
	}

	log.Println("Build completed successfully")
}
