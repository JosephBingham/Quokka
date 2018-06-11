package main

import ("fmt"
	"os"
	"bufio"
)

func getClass(char int ) int {
		if char == 0 {
			return -1
		} else if char <= 32{
			return 0
		} else if char >= 48 && char <= 57{
			return 1
		} else if char >= 65 && char <= 90{
			return 2
		} else if char >= 97 && char <= 122{
			return 3
		} else if char == 58{
			return 4		
		} else {
			panic("this is not something I know, yet")
		}
		
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	class := 0
	token := ""
	for _, char := range text {
		if getClass(int(char)) == 0 {
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

