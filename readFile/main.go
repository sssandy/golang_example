package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 读取文件并按行处理
func readFileAndProcessByLine (path string, proFunc func(string) interface{}) (ret *[]interface{}) {
	ret = &[]interface{}{}

	// 默认后置处理方法
	if proFunc == nil {
		proFunc = func (str string) interface{} {
			// 返回 '\t' 分割的数据
			ret := strings.Split(str, "\t")

			return ret
		}
	}

	fi, err := os.Open(path)
	if err != nil {
		fmt.Printf("FILE ERROR: %s\n", err)
		return ret
	}
	defer fi.Close()

	// 得到一个reader
	br := bufio.NewReader(fi)

	// 按行处理
	for {
		str, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}

		// 后置函数处理
		strRet := (proFunc)(string(str))

		*ret = append(*ret, strRet)
	}

	return ret
}

func main () {
	path := "readFile/data/dict.txt"
	ret := readFileAndProcessByLine(path, nil)

	fmt.Println((*ret)[0])
}
