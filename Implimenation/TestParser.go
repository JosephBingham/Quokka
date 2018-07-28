package main

import ("fmt"
	"os"
	"bufio"
)

func getClass(char int ) int {
		if char == 0 {
			return -1 //EOF, stop running
		} else if char <= 32 || char == 40 || char == 41 {
			return 0 //White space or something I don't care about, ignore
		} else if char >= 48 && char <= 57 {
			return 1 //Numeric literal
		} else if char >= 65 && char <= 90 {
			return 2 //Uppercase letter, type or function
		} else if char >= 97 && char <= 122 {
			return 3 //Lowercase letter, variable
		} else if char == 58 {
			return 4 //Colon 
		} else if char == 124 {
			return 5 //Pipe, for quantum variables 
		} else if char == 61 {
			return 6 //Equals, for assignment
		} else if char == 62 {
			return 7 //Left angled bracket, for quantum types
		} else if char == 59 {
			return 8 //Semicolon, end of statement
		} else if char >= 42 && char <= 47 {
			return 9 //Math operators
		} else if char == 126 {
			return -2
		} else {
			fmt.Println(string(char))
			panic("this is not something I know, yet")
		}
		
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	class := 0
	token := ""
	commentFlag := 0
	for _, char := range text {
		if getClass(int(char)) == -2 {
			commentFlag = (commentFlag + 1) % 2
			fmt.Println(commentFlag)
			continue				
		}
		if getClass(int(char)) == 0 || commentFlag == 1{
			continue
		} else if getClass(int(char)) != class {
			fmt.Println(token)
			token = string(char)
			class = getClass(int(char))
		} else {
			token += string(char)
		}
	}
	fmt.Println(token)
}

