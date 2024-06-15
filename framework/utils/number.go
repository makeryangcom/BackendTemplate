package Utils

import "fmt"

func FormatKUnit(number int) string {
	if number < 1000 {
		return fmt.Sprintf("%d", number)
	}
	kValue := float64(number) / 1000.0
	return fmt.Sprintf("%.1fK", kValue)
}
