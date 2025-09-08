package student
import "github.com/01-edu/z01"
func PrintIf(str string) string {
	if len(str) < 3{
		z01.PrintRune('Invalid Input')

	}else {
		z01.PrintRune('G')
	}
	return 


}