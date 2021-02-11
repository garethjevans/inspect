package cmd

// Logger an interface to allow mocking out of the Println func.
type Logger interface {
	Println(message string)
}
