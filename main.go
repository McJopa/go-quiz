package main
import (
	"fmt"
	"flag"
	"encoding/csv"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "specify csv of questions and answers with format 'question,answer'") 
	time := flag.Int("time", 30, "quiz length time") 
	flag.Parse()
	questions, ans := parseCSV(csvFilename)
	quiz(time, questions, ans)
}

func parseCSV(csvFilename *string) ([]string, []string){
	f, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open csv file: %s\n", *csvFilename) 
		os.Exit(1)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()

	questions := make([]string, len(records))
	answers := make([]string, len(records))
	if err != nil {
		fmt.Printf("Failed to read csv file: %s\n", *csvFilename) 
		os.Exit(1)
	}
	for i, record := range records {
		questions[i] = record[0]
		answers[i] = record[1]
	}
	return questions, answers
}

func quiz (t *int, questions []string, ans []string) {
	timer := time.NewTimer(time.Duration(*t) * time.Second)
	score := 0
	for i := range questions {
		fmt.Println(questions[i])
		aChannel := make(chan string)
		go func (){
			var input string
			fmt.Scanln(&input)
			aChannel <-input
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", score, len(questions))
			return	
		case answer := <-aChannel:
			if answer == ans[i] {
				score++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", score, len(questions))
	return
}

