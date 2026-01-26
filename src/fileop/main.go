package main

import (
	"bufio"
	"fmt"
	"os"
)

// 按行读
func TestReadLine() {
	file, err := os.Open(`.\hello.txt`)
	if err != nil {
		panic(err)		
	}
	defer file.Close()
	buf := bufio.NewScanner(file)
	buf.Split(bufio.ScanWords)
	var index int
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())
	}
}

// 文件写入
func TestWriteFile() {
	file, err := os.OpenFile(`./w.txt`, os.O_RDONLY | os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte("你干嘛"))
}

// 创建目录
func TestCreateDir() {
	const dirname string = "test"
	// 判断目录是否存在
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("目录已经存在")
			return
		} else {
			panic(err)
		}
	}
	fmt.Println("目录创建成功")
}



func main() {
	// byteData, err := os.ReadFile(`./hello.txt`)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(byteData))

	TestReadLine()

	// TestWriteFile()

	TestCreateDir()
}