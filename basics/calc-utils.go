package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func calculateAge(year, month, day int) int {
	var (
		now      = time.Now()
		ageYears = (now.Year() - year) - 1
	)
	if (int(now.Month()) >= month) && (now.Day() >= day) {
		ageYears += 1
	}
	return ageYears
}

func returnManyExample() (int, int, int) {
	return 1, 2, 3
}

func reverseString(str string) string {
	var finalStr strings.Builder
	for lenStr := len(str) - 1; lenStr >= 0; lenStr-- {
		finalStr.WriteByte(str[lenStr])
	}
	return finalStr.String()
}

func binaryToDecimal(bytesString string) int64 {
	reversedString := reverseString(bytesString)
	var sum int64
	for i := 0; i < len(reversedString); i++ {
		localRes, err := strconv.ParseInt(string(reversedString[i]), 10, 64)
		if err != nil {
			return -1
		}
		sum += int64(math.Pow(2, float64(i))) * localRes
	}
	return sum
}

func goWhileLoop(num int) {
	i := 0
	for i < num {
		fmt.Println("current number: ", i)
		i++
	}
}

func goBreakLoops(randomNum int) {
	i := 0
	for {
		fmt.Println("In loop ...", i)
		if i == randomNum {
			break
		}
		i++
	}
}

func diceRoller(dice, sides, rolls int) {
	randomGem := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < dice; i++ {
		currentRoll := rolls
		for currentRoll > 0 {
			randomNum := randomGem.Intn(sides)
			fmt.Println("dice num:", i, "roll num: ", rolls-currentRoll, "result:", randomNum)
			currentRoll--
		}
	}
}
