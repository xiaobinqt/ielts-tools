package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"strings"
)

func Check(ctx *cli.Context) (err error) {
	pathArg := ctx.Args().Get(0) // ylk.8.1.1
	path := strings.Split(pathArg, ".")
	if len(path) < 4 {
		err = errors.Wrapf(err, "路径错误..")
		return err
	}

	dir, chapter, segment, test := path[0], path[1], path[2], path[3]
	exePath, err := os.Executable()
	if err != nil {
		err = errors.Wrapf(err, "获取 Executable 错误")
		return err
	}

	dirwd := filepath.Dir(exePath)

	filePath := fmt.Sprintf("%s%c%s%c%s%c%s%c%s.%s.txt", dirwd, os.PathSeparator, CORPUS, os.PathSeparator,
		dir, os.PathSeparator, chapter, os.PathSeparator, segment, test)

	if strings.HasSuffix(filePath, ".txt.txt") {
		filePath = strings.ReplaceAll(filePath, ".txt.txt", ".txt")
	}

	original, err := readTxt(filePath)
	if err != nil {
		err = errors.Wrapf(err, "读取原始文件出错")
		return err
	}

	// 读取用户听写的单词
	dicPath := fmt.Sprintf("%s%c%s.txt", dirwd, os.PathSeparator, pathArg)
	if strings.HasSuffix(dicPath, ".txt.txt") {
		dicPath = strings.ReplaceAll(dicPath, ".txt.txt", ".txt")
	}
	dicPhrase, err := readTxt(dicPath)
	if err != nil {
		err = errors.Wrapf(err, "读取用户听写的内容出错")
		return err
	}
	//printArray(dicPhrase)

	// 简单填充
	if len(dicPhrase) != len(original) {
		num := len(original) - len(dicPhrase)
		for i := 0; i < num; i++ {
			dicPhrase = append(dicPhrase, "")
		}
	}

	errWords := make([]string, 0)
	for index, o := range original {
		if o != dicPhrase[index] {
			errWords = append(errWords, fmt.Sprintf("%s  |  %s ", o, dicPhrase[index]))
		}
	}

	if len(errWords) == 0 {
		fmt.Println("恭喜你，全部正确，散花!!!!")
		return nil
	}

	printStr := strings.Join(errWords, "\n")
	fmt.Println()
	fmt.Println(fmt.Sprintf("这一小节总共 %d 个单词, 错了 %d 个单词, 正确率为 %d %%",
		len(original), len(errWords), int((float64(len(errWords))/float64(len(original)))*100)))
	fmt.Println()
	fmt.Println(printStr)

	return nil
}

func printArray(arr []string) {
	for _, each := range arr {
		fmt.Println(each)
	}
	fmt.Println("-------------------------------------------")
	if len(arr) > 0 {
		fmt.Println(fmt.Sprintf("%s|", arr[len(arr)-1]))
	}
	fmt.Println("-------------------------------------------")
}

func readTxt(filePath string) (phrases []string, err error) {
	phrases = make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		err = errors.Wrapf(err, "readTxt 读取文件失败")
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	// 逐行读取文件
	for scanner.Scan() {
		line := scanner.Text()

		// 使用逗号分割行
		parts := strings.Split(line, ",")

		// 获取第一个逗号前的单词或短语
		phrase := strings.TrimSpace(parts[0])

		// 如果有多个单词组成的词组，将单词之间的多余空格替换为一个空格
		phrase = strings.Join(strings.Fields(phrase), " ")

		// 添加到结果数组中
		phrases = append(phrases, strings.ToLower(phrase))
	}

	// 检查是否有错误
	if err = scanner.Err(); err != nil {
		err = errors.Wrapf(err, "读取文件失败...2222")
		return nil, err
	}

	return phrases, nil
}
