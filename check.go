package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func Check(ctx *cli.Context) (err error) {
	pathArg := ctx.Args().Get(0) // ylk.11.1
	path := strings.Split(pathArg, ".")
	if len(path) < 3 {
		err = errors.Wrapf(err, "路径错误..")
		return err
	}

	dir, chapter, test := path[0], path[1], path[2]
	dirwd, _ := os.Getwd()
	// 分 2 种方式查找
	filePath := fmt.Sprintf("%s/%s/%s/%s.%s.txt", dirwd, dir, chapter, chapter, test)
	file, err := os.Open(filePath)
	if err != nil {
		err = errors.Wrapf(err, "读取文件失败,%s", pathArg)
		return err
	}

	// 创建一个字符串数组来存储结果
	var result = make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// 使用逗号分隔行
		parts := strings.Split(line, ",")
		if len(parts) > 0 {
			// 获取第一个逗号前的单词或短语
			firstWord := strings.TrimSpace(parts[0])
			// 添加到结果数组中
			result = append(result, firstWord)
		}
	}

	// 检查扫描过程中是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// 打印结果
	fmt.Println(result)

	return nil
}
