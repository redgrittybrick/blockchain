package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	started := time.Now()
	lineCount := 0
	date := ""
	size := 0
	txType := ""

	Total  := make(map[string]int)
	Segwit := make(map[string]int)
	Large  := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineCount = lineCount + 1
		line := scanner.Text()
		f := strings.Split(line, ",")
		switch f[0] {
		case "B": // B,count,version,height,date,txcount,target,id
			dateTime := f[4]
			date = dateTime[:10]
		case "T": // T,count,size,version,type,ins,outs,lock,value,id
			size, _ = strconv.Atoi(strings.TrimSpace(f[2]))
			txType = f[4] // "Legacy" or "Segwit"
			Total[date] = Total[date] + size
			if txType == "Segwit" {
				Segwit[date] = Segwit[date] + size
			}
			if size > 1000 {
				Large[date] = Large[date] + size
			}
		default:
			fmt.Printf("Unexpected value `%s` in 1st field of line %d\n",
				f[0], "\n")
		}

	} // Scan()

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var dates []string
	for k := range Total {
		dates = append(dates, k)
	}
	sort.Strings(dates)
	fmt.Println("Period,Transactions,Segwit,Large")
	for _, d := range dates {
		fmt.Printf("%s, %d, %d, %d\n", d, Total[d], Segwit[d], Large[d])
	}

	fmt.Fprintf(os.Stderr, "Processed %d records in %s\n",
		lineCount, time.Since(started)
	)
}
