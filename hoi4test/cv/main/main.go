package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Province 结构体用于解析JSON数据
type Province struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func main() {
	// 读取provinces.json文件
	data, err := os.ReadFile("provinces.json")
	if err != nil {
		fmt.Printf("读取文件错误: %v\n", err)
		return
	}

	// 解析JSON数据
	var provinces []Province
	err = json.Unmarshal(data, &provinces)
	if err != nil {
		fmt.Printf("解析JSON错误: %v\n", err)
		return
	}

	// 创建输出文件
	outputFile, err := os.Create("cv.txt")
	if err != nil {
		fmt.Printf("创建文件错误: %v\n", err)
		return
	}
	defer outputFile.Close()

	// 处理省份名称并存储
	var processedProvinces []string
	for _, province := range provinces {
		// 去掉后缀：省、市、自治区
		cleanName := province.Name
		cleanName = strings.TrimSuffix(cleanName, "省")
		cleanName = strings.TrimSuffix(cleanName, "市")
		cleanName = strings.TrimSuffix(cleanName, "自治区")

		// 去掉民族名称：回族、壮族、维吾尔等
		cleanName = strings.ReplaceAll(cleanName, "回族", "")
		cleanName = strings.ReplaceAll(cleanName, "壮族", "")
		cleanName = strings.ReplaceAll(cleanName, "维吾尔", "")

		processedProvinces = append(processedProvinces, cleanName)
	}

	// 按照中国革命重要性排序
	// 首先：辽宁、山东、福建（海军重要省份）
	priorityProvinces := []string{"辽宁", "山东", "福建"}

	// 然后按照中国革命历史重要性排序
	revolutionaryOrder := []string{
		"江西",  // 井冈山革命根据地
		"陕西",  // 延安革命圣地
		"湖南",  // 毛泽东故乡，秋收起义
		"湖北",  // 武汉，辛亥革命
		"广东",  // 广州，国民革命
		"浙江",  // 嘉兴南湖，中共一大
		"上海",  // 中共一大召开地
		"北京",  // 首都，五四运动
		"四川",  // 长征重要路线
		"贵州",  // 遵义会议
		"云南",  // 长征路线
		"甘肃",  // 长征终点
		"宁夏",  // 长征路线
		"青海",  // 长征路线
		"西藏",  // 解放西藏
		"新疆",  // 解放新疆
		"内蒙古", // 解放内蒙古
		"黑龙江", // 东北抗联
		"吉林",  // 东北抗联
		"河北",  // 华北抗日
		"山西",  // 晋察冀抗日
		"河南",  // 中原抗日
		"安徽",  // 新四军
		"江苏",  // 新四军
		"天津",  // 华北重要城市
		"重庆",  // 抗战陪都
		"海南",  // 琼崖纵队
		"广西",  // 百色起义
	}

	// 创建最终排序列表
	var finalOrder []string

	// 1. 添加优先省份
	for _, priority := range priorityProvinces {
		finalOrder = append(finalOrder, priority)
	}

	// 2. 添加按革命重要性排序的省份
	for _, revolutionary := range revolutionaryOrder {
		finalOrder = append(finalOrder, revolutionary)
	}

	// 3. 添加其他省份（按原顺序）
	for _, province := range processedProvinces {
		// 检查是否已经在前面添加过
		found := false
		for _, added := range finalOrder {
			if added == province {
				found = true
				break
			}
		}
		if !found {
			finalOrder = append(finalOrder, province)
		}
	}

	// 输出排序后的省份
	for _, province := range finalOrder {
		outputLine := fmt.Sprintf("\t\t\"%s号\"\n", province)
		_, err := outputFile.WriteString(outputLine)
		if err != nil {
			fmt.Printf("写入文件错误: %v\n", err)
			return
		}

		// 同时在控制台显示
		fmt.Print(outputLine)
	}

	// 添加香港、澳门、台湾
	additionalRegions := []string{"香港", "澳门", "台湾"}
	for _, region := range additionalRegions {
		outputLine := fmt.Sprintf("\t\t\"%s号\"\n", region)
		_, err := outputFile.WriteString(outputLine)
		if err != nil {
			fmt.Printf("写入文件错误: %v\n", err)
			return
		}

		// 同时在控制台显示
		fmt.Print(outputLine)
	}

	fmt.Println("结果已保存到 cv.txt 文件中")
}
