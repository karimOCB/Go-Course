package main

// import "fmt"

func getMessageCosts(messages []string) []float64 {
	messagesCosts := make([]float64, len(messages))
	for i := 0; i < len(messagesCosts); i++ {
		messagesCosts[i] = float64(len(messages[i])) * 0.01
	}
	return messagesCosts
}

// don't edit below this line

// Exercise lacking in the original github repository. But copied from the assignment in the youtube video for this course.