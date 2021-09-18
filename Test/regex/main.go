package main

import (
	"log"
	"regexp"
)

var rx *regexp.Regexp = regexp.MustCompile(`\[cap\]`)

func main() {
	str := "奥科吉防护服【adasfafs】afsaf[cap]暗号阿顺"
	findBytes := rx.Find([]byte(str))
	log.Println("Find: ", findBytes)
	findAllBytes := rx.FindAll([]byte(str), -1)
	log.Println("FindAll: ", findAllBytes)
	findAllIdx := rx.FindAllIndex([]byte(str), -1)
	log.Println("FindAllIdx: ", findAllIdx)

	findAllString := rx.FindAllString(str, -1)
	log.Println("FindAllString: ", findAllString)
	findAllStrIdx := rx.FindAllStringIndex(str, -1)
	log.Println("FindAllStringIdx: ", findAllStrIdx)
	findAllStringSub := rx.FindAllStringSubmatch(str, -1)

	log.Println("FindAllStringSub: ", findAllStringSub)
	findAllStrSubIdx := rx.FindAllStringSubmatchIndex(str, -1)
	log.Println("FindAllStringSubmatchIndex", findAllStrSubIdx)

	findAllSub := rx.FindAllSubmatch([]byte(str), -1)
	log.Println("FindAllSubmatch", findAllSub)
	findAllSubIdx := rx.FindAllSubmatchIndex([]byte(str), -1)
	log.Println("FindAllSubmatchIndex", findAllSubIdx)
}
