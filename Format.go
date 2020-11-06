package timeformatter

import (
	"strings"
)

func Format(f string) string {
	if string(f[0]) == "+" {
		f = strings.Replace(f, "+", "", 1)

		f = strings.Replace(f, "%D", "%m/%d/%y", -1)
		f = strings.Replace(f, "%F", "%Y-%m-%d", -1)
		f = strings.Replace(f, "%R", "%H:%M", -1)
		f = strings.Replace(f, "%T", "%H:%M:%S", -1)
		f = strings.Replace(f, "%n", "\n", -1)
		f = strings.Replace(f, "%t", "\t", -1)

		f = strings.Replace(f, "%Y", "2006", -1)
		f = strings.Replace(f, "%y", "06", -1)

		f = strings.Replace(f, "%m", "01", -1)
		f = strings.Replace(f, "%-m", "1", -1)
		f = strings.Replace(f, "%b", "Jan", -1)
		f = strings.Replace(f, "%h", "Jan", -1)
		f = strings.Replace(f, "%B", "January", -1)

		f = strings.Replace(f, "%a", "Mon", -1)
		f = strings.Replace(f, "%A", "Monday", -1)

		f = strings.Replace(f, "%d", "02", -1)
		f = strings.Replace(f, "%-d", "2", -1)
		f = strings.Replace(f, "%e", " 2", -1)
		f = strings.Replace(f, "%j", "002", -1)

		f = strings.Replace(f, "%-H", "15", -1)
		f = strings.Replace(f, "%I", "03", -1)
		f = strings.Replace(f, "%-I", "3", -1)

		f = strings.Replace(f, "%M", "04", -1)
		f = strings.Replace(f, "%-M", "4", -1)

		f = strings.Replace(f, "%S", "05", -1)
		f = strings.Replace(f, "%-s", "5", -1)
		f = strings.Replace(f, ".%N", ".999999999", -1)
		f = strings.Replace(f, "%N", ".999999999", -1)

		for i := 0; i < len(f); i++ {
			j := strings.Index(f[i:len(f)-1], "%")
			if j == -1 {
				break
			}
			if string(f[j+1]) == "%" {
				f = strings.Replace(f, "%%", "%", 1)
			}
		}
		pos := strings.Index(f, "%")
		if pos != -1 {
			panic("Not support format: " + string(f[pos:pos+1]))
		}
	}
	return f
}
