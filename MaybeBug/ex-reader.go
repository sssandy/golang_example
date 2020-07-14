// https://tour.go-zh.org/methods/22
// golang的练习题

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func (r MyReader) Read(b []byte) (int, error) {
	// fmt.Println(&b)
	// b = append(b, 'A') // 这条是往后追加... 肯定不对
	// b[0] = 'A'         // 这种写法有可能传入b 的length == 0 导致错误

	copy(b, []byte{'A'})  // 正确而简介的写法
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
