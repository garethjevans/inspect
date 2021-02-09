package mock

type LoggerMock struct {
	Messages []string
}

func (l *LoggerMock) Println(message string) {
	l.Messages = append(l.Messages, message)
}
