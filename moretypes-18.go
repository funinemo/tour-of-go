package main

import "golang.org/x/tour/pic"

# Pic 画像生成
# https://godoc.org/golang.org/x/tour/pic#example-Show
func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8,dy)
	for i:=0;i<dy;i++ {
		s := make([]uint8,dx)
		for j:=0;j<dx;j++ {
			s[j] = (uint8(j)+uint8(i))/2
		}
		ss[i] = s
	}
	return ss
}

func main() {
	pic.Show(Pic)
}
