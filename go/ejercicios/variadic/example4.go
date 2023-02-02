
// Go program to illustrate the
// concept of variadic function
package main
 
import (
    "fmt"
    "strings"
)
 
// Variadic function to join strings
func joinstr(elements ...string) string {
    return strings.Join(elements, "-")
}
 
func main() {
    // pass a slice in variadic function
    elements := []string{"geeks", "FOR", "geeks"}
	fmt.Println(elements)
    fmt.Println(joinstr(elements...))
}