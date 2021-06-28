package quiz

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type Quiz struct {
	TotalQuestions int
	CorrectAnswers int
}

func (q *Quiz) AddCorrectAnswer() {
	q.CorrectAnswers++
}

func (q *Quiz) AddQuestion() {
	q.TotalQuestions++
}

func ShuffleQuestions(questions [][]string) [][]string {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	return questions
}

func HandleQuiz(q *Quiz, data *csv.Reader, shuffle bool) {
	records, err := data.ReadAll()
	CheckError(err)
	if shuffle {
		records = ShuffleQuestions(records)
	}
	for _, record := range records {

		q.AddQuestion()

		question := record[0]
		answer := record[1]

		fmt.Printf("\nWhat is the result of %s ?\n", question)

		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		if strings.TrimSpace(answer) == strings.TrimSpace(input.Text()) {
			q.AddCorrectAnswer()
		}
	}
}

func ParseCsv(filename string) *csv.Reader {
	fr, err := ioutil.ReadFile(filename)

	CheckError(err)

	filecontents := string(fr)
	return csv.NewReader(strings.NewReader(filecontents))
}

func ParseFlags() (string, int, bool) {
	var filename string
	var timerSeconds int
	var shouldShuffle bool

	flag.StringVar(&filename, "file", "", "Input file that should be read")
	flag.IntVar(&timerSeconds, "timer", 30, "Time after the quiz exits")
	flag.BoolVar(&shouldShuffle, "shuffle", false, "Shuffle questions")
	flag.Parse()

	return filename, timerSeconds, shouldShuffle
}

func StartQuiz(timerSeconds int) {
	fmt.Printf("Press [Enter] to start the quiz. After that you have %d seconds to finish.", timerSeconds)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func StopQuiz(q *Quiz) {
	fmt.Printf("\n\nThere were %d of %d answers correct!!!!\n", q.CorrectAnswers, q.TotalQuestions)
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
