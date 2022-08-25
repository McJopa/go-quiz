package main
import (
	"fmt"
	"flag"
	"encoding/csv"
	"os"
	"log"
)

func main() {
	input := flag.String("i", "problems.csv", "specify csv file for input") 
	flag.Parse()
	
	questions, ans := parseCSV(input)
	quiz(questions, ans)
}

func parseCSV(input *string) ([]string, []string){
	var questions []string
	var ans []string

	f, err := os.Open(*input)

	if err != nil {  
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		questions = append(questions, record[0])
		ans = append(ans, record[1])
	}

	return questions, ans
}

func quiz (questions []string, ans []string) {
	correct := 0
	for i := range questions {
		fmt.Println(questions[i])
		var input string
		fmt.Scanln(&input)
		if input == ans[i] {
			correct += 1
		}
	}
	fmt.Printf("score: %d/%d\n", correct, len(questions))
}
