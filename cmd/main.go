package main

import (
	"os"
	"time"

	q "quiz/internal/quiz"
)

func main() {

	quiz := &q.Quiz{
		TotalQuestions: 0,
		CorrectAnswers: 0,
	}

	filename, timerSeconds, shuffle := q.ParseFlags()

	timer := time.NewTimer(time.Duration(timerSeconds) * time.Second)

	go func() {
		<-timer.C
		q.StopQuiz(quiz)
		os.Exit(1)
	}()

	q.HandleQuiz(quiz, q.ParseCsv(filename), shuffle)

	stop := timer.Stop()

	if stop {
		q.StopQuiz(quiz)
	}
}
