package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const banner = `
 ╔══════════════════════════════════╗
 ║   🎲  Random Content Generator  ║
 ╚══════════════════════════════════╝`

func printSection(title, content string) {
	width := 40
	fmt.Println(strings.Repeat("─", width))
	fmt.Printf(" %-10s │ %s\n", title, content)
}

func printFooter() {
	fmt.Println(strings.Repeat("─", 40))
}

func formatNumber(n int) string {
	return fmt.Sprintf("%d", n)
}

func formatQuote(q string) string {
	return fmt.Sprintf("\"%s\"", q)
}

func formatColor(c RGB) string {
	return fmt.Sprintf("RGB(%d, %d, %d) / %s", c.R, c.G, c.B, c.Hex())
}

func main() {
	mode := flag.String("mode", "all", "output mode: number, quote, color, or all")
	flag.Parse()

	fmt.Println(banner)
	fmt.Println()

	switch *mode {
	case "number":
		printSection("Number", formatNumber(RandomInt(1, 100)))
		printFooter()
	case "quote":
		printSection("Quote", formatQuote(RandomQuote()))
		printFooter()
	case "color":
		printSection("Color", formatColor(RandomColor()))
		printFooter()
	case "all":
		printSection("Number", formatNumber(RandomInt(1, 100)))
		printSection("Quote", formatQuote(RandomQuote()))
		printSection("Color", formatColor(RandomColor()))
		printFooter()
	default:
		fmt.Fprintf(os.Stderr, "unknown mode: %s (use number, quote, color, or all)\n", *mode)
		os.Exit(1)
	}
}
