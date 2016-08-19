package RomanNumeral

func answer() int {
    return 6 * 7
} 

func answer12(i int )string{ 
if i == 1 {return "I" } 
if i == 4 {return "IV" }
if i == 5 {return "V" } 
if i == 9 {return "IX" } 
if i == 10 {return "X" } 
if i == 40 {return "XL" }
if i == 50 {return "L" } 
if i == 90 {return "XC" } 
if i == 100 {return "C" } 
if i == 400 {return "CD" }
if i == 500 {return "C" } 
if i == 900 {return "CM" } 
if i == 1000 {return "M" } 

return "fails"
} 
