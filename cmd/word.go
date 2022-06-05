package cmd

import (
	"log"
	"phalanx/internal/word"
	"strings"

	"github.com/spf13/cobra"
)

// 定义能力
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转为大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转为小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转为下划线
)

// 目标 Word
var str string

// 选择模式
var mode int8

// 描述
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1: 全部转为大写",
	"2: 全部转为小写",
	"3: 下划线转大写驼峰",
	"4: 下划线转小写驼峰",
	"5: 驼峰转下划线",
}, "\n")

// 子命令
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词的格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		}
		log.Printf("输出结果: %s", content)
	},
}

// 初始化需要的参数
func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词的内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
