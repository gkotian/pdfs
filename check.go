package main

import(
    "fmt"
    "os"
    "log"
    "io"
    "encoding/csv"
    // "bufio"
)

func main() {
	csvfile, err := os.Open("output.tsv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	// r := csv.NewReader(bufio.NewReader(csvfile))

    r.Comma = '\t'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
        fmt.Printf("Record Num: %s, Value: %s, Name: %s, City: %s\n", record[0], record[1], record[2], record[3])
	}
}
