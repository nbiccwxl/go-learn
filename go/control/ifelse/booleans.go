package main
import "fmt"
func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}
}

// func init() {
// 	if runtime.GOOS == "windows" {
// 		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")		
// 	} else { //Unix-like
// 		prompt = fmt.Sprintf(prompt, "Ctrl+D")
// 	}
// }