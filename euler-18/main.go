package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var memo = make(map[int]int)

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func convertToInt(data []string) []int {
	var res []int
	for _, v := range data {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln("Error while conversion to int from string", err)
		}
		res = append(res, num)
	}
	return res
}

func findMaxSumPaths(triangle *[][]int, row int, col int) int{
	totalRows := len(*triangle)
	if row >= totalRows {
		return 0
	}
	if val, ok := memo[(*triangle)[row][col]]; ok {
		return val
	}
	sum := (*triangle)[row][col]
	suml := findMaxSumPaths(triangle, row+1, col)
	sumr := findMaxSumPaths(triangle, row+1, col+1)
	sum  += max(suml, sumr)
	return sum
}

func main() {
	data, err := ioutil.ReadFile("triangle.txt")
	if err != nil {
		log.Fatalln(err)
	}
	rows := strings.Split(string(data), "\n")
	var triangle [][]int
	for _, row := range rows {
		xs := strings.Split(row, " ")
		xi := convertToInt(xs)
		triangle = append(triangle, xi)
	}
	maxSum := findMaxSumPaths(&triangle, 0, 0)
	fmt.Println(maxSum)
}
