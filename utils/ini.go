package utils

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var defaultSection = "root"

type ConfigINI struct {
	data map[string]([]ConfigKV)
}

type ConfigKV struct {
	K string
	V string
}

// 取值
func (conf *ConfigINI) Value(section string, key string, defaultValue ...string) string {
	// 章节
	section = strings.TrimSpace(section)
	// 键
	key = strings.TrimSpace(key)
	// 默认值
	if section == "" {
		section = defaultSection
	}
	defaultStr := ""
	if len(defaultValue) > 0 {
		defaultStr = defaultValue[0]
	}
	// 获取对应章节的键值对列表
	sectionKV := conf.data[section]
	for _, kv := range sectionKV {
		// 匹配
		if kv.K == key {
			return kv.V
		}
	}
	return defaultStr
}

// 读取 .ini 配置文件
func LoadINI(fileName string) (*ConfigINI, error) {
	// 打开文件
	file, err1 := os.Open(fileName)
	config := &ConfigINI{}
	if err1 == nil {
		// 安全关闭文件
		defer file.Close()
		scanner := bufio.NewScanner(file)
		// 保存章节[键值对]
		kv := make(map[string]([]ConfigKV))
		// 自上而下读取章节，没有章节的配置必须写在前面
		currentSection := defaultSection
		for scanner.Scan() {
			// 获取每行内容
			line := scanner.Text()
			// 去掉前后空格
			line = strings.TrimSpace(line)
			// 如果不是注释，才处理
			if !strings.HasPrefix(line, ";") {
				// 匹配章节
				regSection, err2 := regexp.Compile(`^\[(.*)\]$`)
				if err2 == nil {
					submatch := regSection.FindStringSubmatch(line) // [0]-匹配的字符串本身，[1]第一个括号的值
					if len(submatch) > 1 {
						section := strings.TrimSpace(submatch[1])
						// 匹配到章节，切换到这个章节
						currentSection = section
					}
				}
				// 匹配键值对
				regKeyValue, err3 := regexp.Compile(`(.*)=(.*)`) // [0]-匹配的字符串本身，[1]key，[2]value
				if err3 == nil {
					submatch := regKeyValue.FindStringSubmatch(line)
					if len(submatch) > 2 {
						// 键
						key := strings.TrimSpace(submatch[1])
						// 值
						value := strings.TrimSpace(submatch[2])
						// 当前章节
						ck := kv[currentSection]
						// 增加
						ck = append(ck, ConfigKV{K: key, V: value})
						kv[currentSection] = ck
					}
				}
			}
		}
		// 保存
		config.data = kv
		return config, nil
	} else {
		return config, err1
	}
}
