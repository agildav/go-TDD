package main

// repeat returns element repeated n times
func repeat(element string, n int) (repeated string) {
	repeated = ""
	for i := 0; i < n; i++ {
		repeated += element
	}
	return
}
