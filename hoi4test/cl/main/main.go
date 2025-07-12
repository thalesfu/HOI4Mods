package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 定义要排除的现代工业园区关键词
var excludeKeywords = []string{
	"工业园区", "经济技术开发区", "文化旅游创意园区", "高新技术产业开发区",
	"经济开发区", "新区", "管理区", "循环化工园区", "芦台经济开发区",
	"汉沽管理区", "海港经济开发区", "冀南新区", "白沟新城", "察北管理区",
	"塞北管理区", "渤海新区", "经济技术开发区",
}

func cleanName(name string) string {
	// 去掉后缀：区、县、市等
	cleanName := name
	cleanName = strings.TrimSuffix(cleanName, "区")
	cleanName = strings.TrimSuffix(cleanName, "县")
	cleanName = strings.TrimSuffix(cleanName, "市")

	// 去掉民族名称
	cleanName = strings.ReplaceAll(cleanName, "满族", "")
	cleanName = strings.ReplaceAll(cleanName, "回族", "")
	cleanName = strings.ReplaceAll(cleanName, "蒙古族", "")
	cleanName = strings.ReplaceAll(cleanName, "土家族", "")
	cleanName = strings.ReplaceAll(cleanName, "苗族", "")
	cleanName = strings.ReplaceAll(cleanName, "侗族", "")
	cleanName = strings.ReplaceAll(cleanName, "布依族", "")
	cleanName = strings.ReplaceAll(cleanName, "哈尼族", "")
	cleanName = strings.ReplaceAll(cleanName, "彝族", "")
	cleanName = strings.ReplaceAll(cleanName, "壮族", "")
	cleanName = strings.ReplaceAll(cleanName, "傣族", "")
	cleanName = strings.ReplaceAll(cleanName, "白族", "")
	cleanName = strings.ReplaceAll(cleanName, "景颇族", "")
	cleanName = strings.ReplaceAll(cleanName, "傈僳族", "")
	cleanName = strings.ReplaceAll(cleanName, "藏族", "")
	cleanName = strings.ReplaceAll(cleanName, "羌族", "")
	cleanName = strings.ReplaceAll(cleanName, "柯尔克孜", "")
	cleanName = strings.ReplaceAll(cleanName, "哈萨克", "")
	cleanName = strings.ReplaceAll(cleanName, "黎族", "")
	cleanName = strings.ReplaceAll(cleanName, "维吾尔", "")
	cleanName = strings.ReplaceAll(cleanName, "朝鲜族", "")

	// 去掉"自治"字样
	cleanName = strings.ReplaceAll(cleanName, "自治", "")

	return cleanName
}

func shouldExclude(name string) bool {
	for _, keyword := range excludeKeywords {
		if strings.Contains(name, keyword) {
			return true
		}
	}
	return false
}

func main() {
	// 读取pca.json文件
	data, err := os.ReadFile("pca.json")
	if err != nil {
		fmt.Printf("读取文件错误: %v\n", err)
		return
	}

	// 解析JSON数据
	var provinces map[string]map[string][]string
	err = json.Unmarshal(data, &provinces)
	if err != nil {
		fmt.Printf("解析JSON错误: %v\n", err)
		return
	}

	// 创建输出文件
	outputFile, err := os.Create("cl.txt")
	if err != nil {
		fmt.Printf("创建cl.txt文件错误: %v\n", err)
		return
	}
	defer outputFile.Close()

	// 收集所有区县
	var allCounties []string

	// 处理所有省份的区县
	for _, cities := range provinces {
		for _, counties := range cities {
			for _, county := range counties {
				// 排除现代工业园区等
				if shouldExclude(county) {
					continue
				}

				cleanCountyName := cleanName(county)

				// 跳过空名称
				if cleanCountyName == "" {
					continue
				}

				allCounties = append(allCounties, cleanCountyName)
				fmt.Printf("处理区县: %s -> %s\n", county, cleanCountyName)
			}
		}
	}

	// 按照类似cv的规则排序
	allCounties = sortCounties(allCounties)

	// 写入文件
	for _, county := range allCounties {
		outputLine := fmt.Sprintf("\t\t\"%s舰\"\n", county)
		_, err := outputFile.WriteString(outputLine)
		if err != nil {
			fmt.Printf("写入文件错误: %v\n", err)
			return
		}
	}

	fmt.Printf("结果已保存到 cl.txt 文件中，共处理了 %d 个区县\n", len(allCounties))
}

// 按照类似cv的规则排序区县
func sortCounties(counties []string) []string {
	// 暂时按原顺序返回，后续可以根据需要实现更复杂的排序
	return counties
}
