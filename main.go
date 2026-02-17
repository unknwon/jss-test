package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.String("mode", "all", "output mode: number, quote, color, or all")
	flag.Parse()

	switch *mode {
	case "number":
		fmt.Printf("Random number (1-100): %d\n", RandomInt(1, 100))
	case "quote":
		fmt.Printf("Random quote: %s\n", RandomQuote())
	case "color":
		c := RandomColor()
		fmt.Printf("Random color: RGB(%d, %d, %d) %s\n", c.R, c.G, c.B, c.Hex())
	case "all":
		fmt.Printf("Random number (1-100): %d\n", RandomInt(1, 100))
		fmt.Printf("Random quote: %s\n", RandomQuote())
		c := RandomColor()
		fmt.Printf("Random color: RGB(%d, %d, %d) %s\n", c.R, c.G, c.B, c.Hex())
	default:
		fmt.Fprintf(os.Stderr, "unknown mode: %s (use number, quote, color, or all)\n", *mode)
		os.Exit(1)
	}
}
