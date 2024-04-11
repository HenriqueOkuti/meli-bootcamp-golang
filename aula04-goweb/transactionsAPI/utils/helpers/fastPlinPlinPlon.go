package fastPlinPlinPlon

import (
	"math/rand"
	"strings"
)

func FastPlinPlinPlon() string {
	plinplinPlons := rand.Intn(10000000)
	var name string
	var sb strings.Builder
	for i := 0; i < plinplinPlons; i++ {
		if i%2 == 0 {
			sb.WriteString("plinplin")
		} else {
			sb.WriteString("plon")
		}
	}
	name = sb.String()
	return name
}
