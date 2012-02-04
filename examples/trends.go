package main

import "fmt"
import "time"
import "github.com/laktek/stack-on-go"

func main() {
	tags := []string{"Go", "Erlang", "Ruby", "Python", "JavaScript"}

	fmt.Printf("No.of Questions for Technology (Last 30 days):\n")

	from_date := time.LocalTime().Seconds() - (60 * 60 * 24 * 30)

	session := stackongo.NewSession("stackoverflow")

	for _, tag := range tags {
		results, err := session.AllQuestions(map[string]string{"tagged": tag, "filter": "total", "fromdate": fmt.Sprintf("%v", from_date)})

		if err != nil {
			fmt.Printf(err.String())
		}

		fmt.Printf("%v  %v\n", tag, results.Total)

	}
}