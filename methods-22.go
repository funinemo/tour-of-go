package main

import "golang.org/x/tour/reader"

// MyReader 型
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// 参照：https://qiita.com/rock619/items/f412d1f870a022c142d0
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
