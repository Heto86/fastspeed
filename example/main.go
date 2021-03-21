package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/Heto86/fastspeed"
)

func main() {
	//fastspeed.GetMeasurments(fastspeed.FastcomType)
	old_stdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout = writer
	scanner := bufio.NewScanner(reader)
	// copy the output in a separate goroutine so printing can't block indefinitely
	var outputs []string
	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			outputs = append(outputs, line)
		}
	}()
	fastspeed.GetMeasurments(fastspeed.FastcomType)
	writer.Close()
	os.Stdout = old_stdout // restoring the real stdout
	for _, output := range outputs {
		fmt.Println(output)
		re_output := regexp.MustCompile(`(Upload|Download): \d{2,3}\.\d{2}`)
		if !re_output.MatchString(output) {
			panic(err)
		}
	}
}
