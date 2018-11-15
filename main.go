package main

import (
	"fmt"
	"os"
	"strings"
	"errors"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Expecting date as argument")
		os.Exit(1)
	}

	cleanDate, err := Format(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cleanDate)
	}
}

func Format(input string) (string, error) {
	var outputString string
	var outputTime time.Time
	var err error

	if strings.Contains(input, "ago") {
		return "", errors.New("todo")
	} else {
		//var splitInput []string;
		var fieldCount int
		var sep string
		if strings.ContainsRune(input, '/') {
			fieldCount = len(strings.Split(input, "/"))
			sep = "/"
		} else if strings.ContainsRune(input, '-') {
			fieldCount = len(strings.Split(input, "-"))
			sep = "-"
		}

		// I'm expecting the format mm-dd-yy or mm-dd-yyyy
		// or maybe mm-dd
		// fuck europeans

		if fieldCount == 2 {
			format := "1" + sep + "2" + sep + "2006"
			outputTime, err = time.Parse(format, input + sep + time.Now().Format("2006"))
			if err != nil {
				return "", err
			}
		} else if fieldCount == 3 {
			if (len(strings.Split(input, sep)[2]) == 4) {
				format := "1" + sep + "2" + sep + "2006"
				outputTime, err = time.Parse(format, input)
			} else {
				format := "1" + sep + "2" + sep + "06"
				outputTime, err = time.Parse(format, input)
			}
		} else {
			return "", errors.New("Expected mm/dd or mm/dd/yy")
		}

		// finally, format as string
		outputString = outputTime.Format("02-01-06")


		return outputString, nil
	}
}

