package example

import "fmt"

func ExampleDemo_DemoHello() {

}

func ExampleHello() {
	greeting := Hello("Jon")
	fmt.Println(greeting)

	// Output: Hello, Jon
}

func ExamplePage() {
	checkIns := map[string]bool {
		"Bob": true,
		"Alice": false,
		"Eve": false,
		"Janet": false,
		"Susan": true,
		"John": false,
	}

	Page(checkIns)

	// Unordered output:
	// Paging Bob; please see the front desk to check in.
	// Paging Susan; please see the front desk to check in.	
}