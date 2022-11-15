package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	LuckyTicketsFor6DigitsNumber()
	tests()
}

func tests() {
	TestLuckyTickets("LuckyTickets")
	TestIterationExponentiation()
}

func TestLuckyTickets(testName string) {
	luckyTickets := LuckyTickets{}
	for i := 0; i <= 9; i++ {
		N, _ := strconv.Atoi(readFromFile("test_data/" + testName + "/test." + strconv.Itoa(i) + ".in"))
		result := luckyTickets.Test(N)
		validValue, _ := strconv.Atoi(readFromFile("test_data/" + testName + "/test." + strconv.Itoa(i) + ".out"))

		fmt.Println(i, ": ", result == validValue)
	}
}

func TestIterationExponentiation() {
	e := Exponentiation{}
	fmt.Println(e.Test(2, 8) == 256)
}

func readFromFile(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	return wr.String()
}
