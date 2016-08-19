package hiker

import ("testing")

func Test_life_the_universe_and_everything(t *testing.T) {
    if (answer() != 42) {
        t.Error("answer() != 42 as expected.")
    }
}


func Test_life_the_universe_and_everything1(t *testing.T) {
    if (answer12(1) != "I") {
        t.Error("fail")
    }
}  

func Test4(t *testing.T) {
    if (answer12(4) != "IV") {
        t.Error("fail")
    }
} 

func Test5(t *testing.T) {
    if (answer12(5) != "V") {
        t.Error("fail")
    }
} 

func Test9(t *testing.T) {
    if (answer12(9) != "IX") {
        t.Error("fail")
    }
} 

func Test10(t *testing.T) {
    if (answer12(10) != "X") {
        t.Error("fail")
    }
}  

func Test40(t *testing.T) {
    if (answer12(40) != "XL") {
        t.Error("fail")
    }
}   

func Test50(t *testing.T) {
    if (answer12(50) != "L") {
        t.Error("fail")
    }
}  

func Test90(t *testing.T) {
    if (answer12(90) != "XC") {
        t.Error("fail")
    }
} 

func Test100(t *testing.T) {
    if (answer12(100) != "C") {
        t.Error("fail")
    }
} 

func Test400(t *testing.T) {
    if (answer12(400) != "CD") {
        t.Error("fail")
    }
} 

func Test500(t *testing.T) {
    if (answer12(500) != "C") {
        t.Error("fail")
    }
} 

func Test900(t *testing.T) {
    if (answer12(900) != "CM") {
        t.Error("fail")
    }
} 

func Test1000(t *testing.T) {
    if (answer12(1000) != "M") {
        t.Error("fail")
    }
}  
