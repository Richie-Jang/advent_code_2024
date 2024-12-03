package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

const (
	input = "input.txt"
)

func checkLine(s string) bool {
	ss := strings.Split(s, " ")
	ar := lo.Map(ss, func(item string, _ int) int {
		return lo.Must1(strconv.Atoi(item))
	})

	checking := func(i int) (isOk bool, dir int) {
		isOk = true
		v := ar[i] - ar[i+1]
		gap := lo.IfF(v < 0, func() int { return v * -1 }).ElseF(func() int { return v })
		if gap == 0 || gap >= 4 {
			isOk = false
			return
		}
		if v < 0 {
			dir = -1
		} else {
			dir = 1
		}
		return
	}

	ok1, d1 := checking(0)
	if !ok1 {
		return false
	}
	for i := 1; i < len(ar)-1; i++ {
		ok2, d2 := checking(i)
		if !ok2 {
			return false
		}
		if d1 != d2 {
			return false
		}
	}
	return true
}

func main() {
	f, er := os.Open(input)
	if er != nil {
		log.Fatal(er)
	}
	defer f.Close()
	br := bufio.NewScanner(f)
	res := 0
	for br.Scan() {
		t := br.Text()
		if t == "" {
			continue
		}
		if checkLine(t) {
			res++
		}
	}

	fmt.Println(res)
}
