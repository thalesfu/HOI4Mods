package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 定义重要城市（不超过80个）
var importantCities = map[string]bool{
	// 直辖市
	"北京": true, "天津": true, "上海": true, "重庆": true,

	// 省会城市
	"石家庄": true, "太原": true, "呼和浩特": true, "沈阳": true, "长春": true, "哈尔滨": true,
	"南京": true, "杭州": true, "合肥": true, "福州": true, "南昌": true, "济南": true,
	"郑州": true, "武汉": true, "长沙": true, "广州": true, "南宁": true, "海口": true,
	"成都": true, "贵阳": true, "昆明": true, "拉萨": true, "西安": true, "兰州": true,
	"西宁": true, "银川": true, "乌鲁木齐": true,

	// 重要经济城市
	"大连": true, "青岛": true, "厦门": true, "深圳": true, "珠海": true,
	"宁波": true, "苏州": true, "无锡": true, "常州": true,
	"佛山": true, "东莞": true, "中山": true,

	// 革命意义重要城市
	"井冈山": true, "瑞金": true, "延安": true, "遵义": true, "嘉兴": true,
	"韶山": true, "古田": true, "百色": true, "西柏坡": true,

	// 重要工业城市
	"唐山": true, "包头": true, "大庆": true, "鞍山": true, "抚顺": true,
	"吉林": true, "齐齐哈尔": true, "徐州": true, "温州": true, "泉州": true,
	"烟台": true, "洛阳": true, "宜昌": true, "株洲": true, "柳州": true,
	"桂林": true, "三亚": true, "绵阳": true, "宝鸡": true, "克拉玛依": true,
}

func cleanName(name string) string {
	// 去掉后缀：市、区、自治州、地区、林区、县等
	cleanName := name
	cleanName = strings.TrimSuffix(cleanName, "市")
	cleanName = strings.TrimSuffix(cleanName, "区")
	cleanName = strings.TrimSuffix(cleanName, "自治州")
	cleanName = strings.TrimSuffix(cleanName, "地区")
	cleanName = strings.TrimSuffix(cleanName, "林区")
	cleanName = strings.TrimSuffix(cleanName, "县")
	cleanName = strings.TrimSuffix(cleanName, "盟")

	// 去掉民族名称
	cleanName = strings.ReplaceAll(cleanName, "土家族苗族", "")
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
	cleanName = strings.ReplaceAll(cleanName, "回族", "")
	cleanName = strings.ReplaceAll(cleanName, "蒙古", "")
	cleanName = strings.ReplaceAll(cleanName, "柯尔克孜", "")
	cleanName = strings.ReplaceAll(cleanName, "哈萨克", "")
	cleanName = strings.ReplaceAll(cleanName, "黎族", "")
	cleanName = strings.ReplaceAll(cleanName, "维吾尔", "")
	cleanName = strings.ReplaceAll(cleanName, "朝鲜族", "")

	// 去掉"自治"字样
	cleanName = strings.ReplaceAll(cleanName, "自治", "")

	return cleanName
}

func main() {
	// 读取pc.json文件
	data, err := os.ReadFile("pc.json")
	if err != nil {
		fmt.Printf("读取文件错误: %v\n", err)
		return
	}

	// 解析JSON数据
	var provinces map[string][]string
	err = json.Unmarshal(data, &provinces)
	if err != nil {
		fmt.Printf("解析JSON错误: %v\n", err)
		return
	}

	// 收集所有城市
	var bcCities []string
	var caCities []string

	// 处理所有城市
	for _, cities := range provinces {
		for _, city := range cities {
			cleanCityName := cleanName(city)

			// 判断是否为重要城市
			if importantCities[cleanCityName] {
				bcCities = append(bcCities, cleanCityName)
				fmt.Printf("重要城市: %s\n", cleanCityName)
			} else {
				caCities = append(caCities, cleanCityName)
				fmt.Printf("普通城市: %s\n", cleanCityName)
			}
		}
	}

	// 按照中国革命重要性排序bc城市
	bcCities = sortBCCities(bcCities)

	// 按照省份重要性排序ca城市
	caCities = sortCACities(caCities)

	// 创建输出文件
	bcFile, err := os.Create("bc.txt")
	if err != nil {
		fmt.Printf("创建bc.txt文件错误: %v\n", err)
		return
	}
	defer bcFile.Close()

	caFile, err := os.Create("ca.txt")
	if err != nil {
		fmt.Printf("创建ca.txt文件错误: %v\n", err)
		return
	}
	defer caFile.Close()

	// 写入bc.txt
	for _, city := range bcCities {
		outputLine := fmt.Sprintf("\t\t\"%s舰\"\n", city)
		_, err := bcFile.WriteString(outputLine)
		if err != nil {
			fmt.Printf("写入bc.txt错误: %v\n", err)
			return
		}
	}

	// 写入ca.txt
	for _, city := range caCities {
		outputLine := fmt.Sprintf("\t\t\"%s舰\"\n", city)
		_, err := caFile.WriteString(outputLine)
		if err != nil {
			fmt.Printf("写入ca.txt错误: %v\n", err)
			return
		}
	}

	fmt.Println("结果已保存到 bc.txt 和 ca.txt 文件中")
}

// 按照中国革命重要性排序bc城市
func sortBCCities(cities []string) []string {
	// 首先：辽宁、山东、福建的重要城市（海军重要省份）
	priorityCities := []string{"大连", "青岛", "厦门", "福州", "沈阳", "济南"}

	// 然后按照中国革命历史重要性排序
	revolutionaryOrder := []string{
		"南昌",   // 南昌起义
		"井冈山",  // 井冈山革命根据地
		"延安",   // 延安革命圣地
		"长沙",   // 毛泽东故乡，秋收起义
		"武汉",   // 武汉，辛亥革命
		"广州",   // 广州，国民革命
		"嘉兴",   // 嘉兴南湖，中共一大
		"上海",   // 中共一大召开地
		"北京",   // 首都，五四运动
		"成都",   // 长征重要路线
		"遵义",   // 遵义会议
		"昆明",   // 长征路线
		"兰州",   // 长征终点
		"银川",   // 长征路线
		"西宁",   // 长征路线
		"拉萨",   // 解放西藏
		"乌鲁木齐", // 解放新疆
		"呼和浩特", // 解放内蒙古
		"哈尔滨",  // 东北抗联
		"长春",   // 东北抗联
		"石家庄",  // 华北抗日
		"太原",   // 晋察冀抗日
		"郑州",   // 中原抗日
		"合肥",   // 新四军
		"南京",   // 新四军
		"天津",   // 华北重要城市
		"重庆",   // 抗战陪都
		"海口",   // 琼崖纵队
		"南宁",   // 百色起义
		"百色",   // 百色起义
	}

	// 创建最终排序列表
	var finalOrder []string

	// 1. 添加优先城市
	for _, priority := range priorityCities {
		for _, city := range cities {
			if city == priority {
				finalOrder = append(finalOrder, city)
				break
			}
		}
	}

	// 2. 添加按革命重要性排序的城市
	for _, revolutionary := range revolutionaryOrder {
		for _, city := range cities {
			if city == revolutionary {
				finalOrder = append(finalOrder, city)
				break
			}
		}
	}

	// 3. 添加其他重要城市（按原顺序）
	for _, city := range cities {
		// 检查是否已经在前面添加过
		found := false
		for _, added := range finalOrder {
			if added == city {
				found = true
				break
			}
		}
		if !found {
			finalOrder = append(finalOrder, city)
		}
	}

	return finalOrder
}

// 按照省份重要性排序ca城市
func sortCACities(cities []string) []string {
	// 暂时按原顺序返回，后续可以根据需要实现更复杂的排序
	return cities
}
