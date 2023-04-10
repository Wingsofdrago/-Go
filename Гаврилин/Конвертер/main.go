package main

import "fmt"


func romanToArabic(s string) int {
  
    romanValues := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

   
    arabic := 0
    prev := 0
    for i := len(s) - 1; i >= 0; i-- {
      
        cur := romanValues[rune(s[i])]

      
        if prev <= cur {
            arabic += cur
        } else { 
            arabic -= cur
        }
        prev = cur
    }

   
    return arabic
}

func main() {
    
    romanNumeral := "XIVXXIVMC"
    arabicNumeral := romanToArabic(romanNumeral)
    fmt.Printf("%s in Arabic numerals is %d\n", romanNumeral, arabicNumeral)
}