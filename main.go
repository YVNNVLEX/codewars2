package main

import (
	"fmt"
	"slices"
	"strings"
)

func main(){
	transf("bonjour_toi")
}
func transf(a string) []rune{
	var chars []rune
	for _, c := range a {
		chars = append(chars, c)
	}
	fmt.Printf("%c ", chars)
	if slices.Contains(chars, rune('-')) || slices.Contains(chars,'_'){
		if slices.Contains(chars, rune('-')){
			pos := slices.Index(chars, rune('-')) + 1
			
			fmt.Print(chars[pos])
		} else if slices.Contains(chars,'_'){
			indexA := slices.Index(chars, rune('_'))-1
			indexB := slices.Index(chars, rune('_'))+1
			posA := string(chars[indexA])
			posB := string(chars[indexB])
			var stringCapA = strings.ToUpper(posA)
			var stringCapB= strings.ToUpper(posB)
			slices.Replace(chars,indexA,indexB,stringCapA,stringCapB)
			fmt.Print(stringCapB, stringCapA)
		}
		
	}
	return chars
}