package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
)

func testReflect() {
	myVal := 43324
	log.Println("reflect val: ", reflect.ValueOf(myVal))
	log.Println("Type of myVal: ", reflect.TypeOf(myVal))
}

func testRegx() {
	fmt.Println("Testing regular expressions...")
	strVal := "I dont know kkkx, ffff15969527725.kxyzwangenglishkl@12wx6.com www.goog45le.com"
	// pattern := "w[a-z]{1,}@[a-z|0-9]{3,}.[a-z]{2,4}"
	pattern := `www.[a-z|0-9]{1,}.{1,}[com|net|cn]`
	matched, err := regexp.MatchString(pattern, strVal)
	if err != nil {
		log.Println("Matching string error.")
		return
	}
	log.Println("Matched: ", matched)
	matchedStr := regexp.MustCompile(pattern).FindString(strVal)
	log.Println("matched string: ", matchedStr)
	fmt.Println("Test regular expressions complete.")
}

func testBits() {
	num := 1 << 10
	fmt.Println(num)
}

func main() {
	fmt.Println("Reflection test.")
	testReflect()
	testRegx()
	testBits()
}
