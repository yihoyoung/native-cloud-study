package main

import "fmt"

func ChangeArray() [5]string {
	// var arrTxt = [5]string{"I", "am", "smart", "and", "strong"}
	var arrTxt = [5]string{"I", "am", "stupid", "and", "weak"}

	for i, _ := range arrTxt {
		if arrTxt[i] == "stupid" {
			arrTxt[i] = "smart"
		}
		if arrTxt[i] == "weak" {
			arrTxt[i] = "strong"
		}
	}

	return arrTxt
}

func main() {
	result := ChangeArray()
	fmt.Println(result)
}
