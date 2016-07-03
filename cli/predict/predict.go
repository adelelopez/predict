package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/adelelopez/predict/api"
	"github.com/adelelopez/predict/src"
	"github.com/fatih/color"
)

func spaces(n int) string {
	ret := ""
	for i := 0; i != n; i++ {
		ret += " "
	}
	return ret
}

func printBucket(b api.Bucket) {
	star := int(b.Star * 100)
	left := int(b.LeftBound * 100)
	right := int(b.RightBound * 100)
	mean := int(b.Mean * 100)

	if left == star {
		left--
	}
	if right == star {
		right++
	}

	blue := color.New(color.FgRed, color.Bold).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	str := red(fmt.Sprintf("%2d", mean)) + "%:" + spaces(left) + blue("[") + spaces(star-left-1) + green("*") + spaces(right-star-1) + blue("]")
	fmt.Println(str)
}

func printPrediction(prediction api.Prediction) {
	outcome := "-"
	createdAt := ""
	probability := 0.0
	if prediction.CreatedAt != nil {
		createdAt = prediction.CreatedAt.Format("01/02/2006 3:04pm")
	}
	if prediction.Outcome != nil {
		outcome = fmt.Sprint(*prediction.Outcome)
	}
	if prediction.Probability != nil {
		probability = *prediction.Probability
	}
	fmt.Printf("%s: %-50s \t%.2f %s\n",
		createdAt,
		prediction.Name,
		probability,
		outcome)
}

var tag string

func init() {
	flag.StringVar(&tag, "tag", "debug", "tag a prediction")
	flag.StringVar(&tag, "t", "debug", "tag a prediction")
}

func main() {
	fs := filestorage.FileStorage{Filename: "~/predictions.txt"}

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Not enough arguments")
		return
	}

	command := flag.Arg(0)

	switch command {
	case "stats":
		stats, err := api.GetStats(&fs)
		stats = new(api.Statistics)
		if stats == nil || err != nil {
			fmt.Println("No statistics available.")
			break
		}
		stats.Buckets = append(stats.Buckets, api.Bucket{
			LeftBound:  0.4,
			RightBound: 0.7,
			Star:       0.5,
			Mean:       0.5,
		})
		fmt.Println("::::=1=3=5==10===15===20===25=======33===============50==============66======75===80===85===90===95===99=")
		for _, b := range stats.Buckets {
			printBucket(b)
		}
		fmt.Println("::::=1=3=5==10===15===20===25=======33===============50==============66======75===80===85===90===95===99=")
		break
	case "hist", "history":
		ps, err := api.GetHistory(&fs)
		if err != nil {
			fmt.Println(err)
			break
		}
		for _, prediction := range ps {
			printPrediction(prediction)
		}
		break
	case "get":
		prediction, _ := fs.GetPrediction("q")
		printPrediction(*prediction)

		break
	case "judge":
		if len(flag.Args()) == 2 {
			outcomeStr := flag.Arg(1)
			outcome := false
			switch outcomeStr {
			case "true", "True", "TRUE", "t", "T":
				outcome = true
				break
			case "false", "False", "FALSE", "f", "F":
				outcome = false
				break
			}

			api.JudgeLastPrediction(outcome, &fs)
		}
		break
	default:
		if len(flag.Args()) == 2 {
			probability := flag.Arg(1)

			probFloat, err := strconv.ParseFloat(probability, 64)

			// TODO: check bounds
			// TODO: decide how to deal with probability and percentages and 1% in particular

			p := api.Prediction{
				Name:        command,
				Probability: &probFloat,
				Tags:        []string{tag},
			}

			_, err = api.CreatePrediction(p, &fs)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: no probability entered.")
		}
	}
}
