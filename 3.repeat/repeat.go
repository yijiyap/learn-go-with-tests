package repeat

import "strings"

func Repeat(ch string, count int) string {
	var ans strings.Builder
	for i := 0; i < count; i++ {
		ans.WriteString(ch)
	}
	return ans.String()
}
