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
	fmt.Println("   kB:", comma(kB))    // "   kB: 1,000"
	fmt.Println("   MB:", comma(MB))    // "   MB: 1,000,000"
	fmt.Println("   GB:", comma(GB))    // "   GB: 1,000,000,000"
	fmt.Println("   TB:", comma(TB))    // "   TB: 1,000,000,000,000"
	fmt.Println("   PB:", comma(PB))    // "   PB: 1,000,000,000,000,000"
	fmt.Println("   EB:", comma(EB))    // "   EB: 1,000,000,000,000,000,000"
	fmt.Println("YB/ZB:", comma(YB/ZB)) // "YB/ZB: 1,000"
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
