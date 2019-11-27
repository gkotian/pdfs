package main

import (
	"fmt"
	"os"

	"github.com/ledongthuc/pdf"
)

var suspiciousRecords []string

func main() {
	_, err := readPdf("combined.pdf")
	if err != nil {
		panic(err)
	}
	return
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}

    of, err := os.Create("output.tsv")
	defer of.Close()
	if err != nil {
        fmt.Println("Failed to create output file")
		return "", err
	}

	totalPage := r.NumPage()
    recordNumber := 0

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		var lastTextStyle pdf.Text
		texts := p.Content().Text
        record := ""
        var skipChars int
        var last_x float64
		for _, text := range texts {
			if isSameSentence(text, lastTextStyle) {
				lastTextStyle.S = lastTextStyle.S + text.S
			} else {
				// fmt.Printf("Font: %s, Font-size: %f, x: %f, y: %f, content: %s \n", lastTextStyle.Font, lastTextStyle.FontSize, lastTextStyle.X, lastTextStyle.Y, lastTextStyle.S)
                if lastTextStyle.Font == "ArialMT" && (lastTextStyle.FontSize > 10) && (lastTextStyle.FontSize < 11) {
                    x := lastTextStyle.X
                    if x > 40 && x < 50 {
                        record = record + "\n"
                        saveRecord(record, of)

                        recordNumber++
                        record = ""

                        if recordNumber <= 9 {
                            skipChars = 1
                        } else if recordNumber <= 100 {
                            skipChars = 2
                        } else if recordNumber <= 1000 {
                            skipChars = 3
                        }
                        record = lastTextStyle.S[:skipChars] + "\t" + lastTextStyle.S[skipChars:]
                        // record = lastTextStyle.S
                    } else {
                        // fmt.Printf("x = %f, last_x = %f ", x, last_x)
                        if (x - last_x < 1) {
                            record = record + " " + lastTextStyle.S
                        } else {
                            // fmt.Printf("Got a new column: %s\n", lastTextStyle.S)
                            record = record + "\t" + lastTextStyle.S
                        }
                    }
                    last_x = x
                }
				lastTextStyle = text
			}
		}
        record = record + "\n"
        saveRecord(record, of)
	}

    fmt.Println("Output saved to 'output.tsv'")
    fmt.Printf("The following %d records need to be checked manually:\n    ", len(suspiciousRecords))
    fmt.Println(suspiciousRecords)
    fmt.Println("Also make sure to manually check the last line of the last record in each page of the input file")
    fmt.Println("Use check.go after manual fixing")
	return "", nil
}

func saveRecord(record string, f *os.File) {
    if len(record) <= 1 {
        return
    }

    recordNum := ""
    numTabs := 0
    for i, c := range record {
        if c == '\t' {
            numTabs++
            if recordNum == "" {
                recordNum = record[:i]
            }
        }
    }

    const expectedNumTabs = 4
    if numTabs != expectedNumTabs {
        suspiciousRecords = append(suspiciousRecords, recordNum)
    }

    // fmt.Println(record)
    // fmt.Println("-------------------------------------------------")
    // f, err := os.Create("/tmp/dat2")
    _, err := f.WriteString(record)
    if err != nil {
        fmt.Println("failed to write record", recordNum)
    }
}

func isSameSentence(t1, t2 pdf.Text) bool {
       // if Y axis changes new line else same line
	if t1.Y != t2.Y {
		return false
	}
	return true
}
