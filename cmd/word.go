package cmd

import (
	"log"
	"strings"

	"felix.bs.com/felix/BeStrongerInGO/01_CLI/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"該子指令支援各種單字格式轉換, 模式如下:",
	"1: 全部單字轉為大寫",
	"2: 全部單字轉為小寫",
	"3: 底線單字轉為大寫駝峰單字",
	"4: 底線單字轉為小寫駝峰單字",
	"5: 駝峰單字轉為底線單字",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "單字格式轉換",
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
		default:
			log.Fatalf("暫不支援該轉換模式, 請執行help word檢視說明文件")
		}

		log.Printf("輸出結果: %s\n", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "請輸入單字內容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "請輸入單字轉換的模式")
}
