// Write const declarations for KB, MB, up through YB as compactly as you can.
package main

import (
	"bytes"
	"fmt"
)

const (
	kB = 1000
	MB = kB * kB
	GB = MB * kB
	TB = GB * kB
	PB = TB * kB
	EB = PB * kB
	ZB = EB * kB
	YB = ZB * kB
)

func main() {
	fmt.Println("   kB:", comma(kB))
	fmt.Println("   MB:", comma(MB))
	fmt.Println("   GB:", comma(GB))
	fmt.Println("   TB:", comma(TB))
	fmt.Println("   PB:", comma(PB))
	fmt.Println("   EB:", comma(EB))
	fmt.Println("YB/ZB:", comma(YB/ZB))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(n int) string {
	var buf bytes.Buffer
	s := fmt.Sprintf("%d", n)
	i := (len(s)-1)%3 + 1

	buf.WriteString(s[:i])
	for ; i < len(s); i += 3 {
		buf.WriteString("," + s[i:i+3])
	}
	return buf.String()
}
