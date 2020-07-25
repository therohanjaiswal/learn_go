package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		locs                       []string
		sizes, beds, baths, prices []int
	)

	rows := strings.Split(data, "\n")

	for _, row := range rows {
		cols := strings.Split(row, separator)

		locs = append(locs, cols[0])

		n, _ := strconv.Atoi(cols[1])
		sizes = append(sizes, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		prices = append(prices, n)
	}

	for _, h := range strings.Split(header, separator) {
		fmt.Printf("%-15s", h)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range rows {
		fmt.Printf("%-15s", locs[i])
		fmt.Printf("%-15d", sizes[i])
		fmt.Printf("%-15d", beds[i])
		fmt.Printf("%-15d", baths[i])
		fmt.Printf("%-15d", prices[i])
		fmt.Println()
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	var (
		sizeAvg  float64
		bedsAvg  float64
		bathsAvg float64
		priceAvg float64
	)

	for i := range rows {
		sizeAvg += float64(sizes[i])
		bedsAvg += float64(beds[i])
		bathsAvg += float64(baths[i])
		priceAvg += float64(prices[i])
	}

	fmt.Printf("%-15s", "")
	fmt.Printf("%-15f", sizeAvg/float64(len(sizes)))
	fmt.Printf("%-15f", bedsAvg/float64(len(beds)))
	fmt.Printf("%-15f", bathsAvg/float64(len(baths)))
	fmt.Printf("%-15f", priceAvg/float64(len(prices)))

}
