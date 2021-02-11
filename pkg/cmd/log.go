package cmd

// Logs an interface to allow mocking out of the Println func.
type Logs interface {
	Println(message string)
}
