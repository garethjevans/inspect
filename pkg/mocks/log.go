package mocks

// LoggerMock the mock logger.
type LoggerMock struct {
	Messages []string
}

// Println the mocked out Println func.
func (l *LoggerMock) Println(message string) {
	l.Messages = append(l.Messages, message)
}
