package srv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zhexiao/edu-parser-rpc/base"
	"github.com/zhexiao/office-parser/excel"
)

type CT_Excel struct {
}

func NewCT_Excel() *CT_Excel {
	return &CT_Excel{}
}

//读取 excel paper 数据
func (c *CT_Excel) paper(fileData []byte) (string, error) {
	data := excel.ConvertFromData(bytes.NewReader(fileData), int64(len(fileData)), "paper")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", base.NewEpError(1, fmt.Sprintf("Excel试卷 json转换失败: %s", err))
	}

	return string(jsonBytes), nil
}

//读取 excel outline 数据
func (c *CT_Excel) outline(fileData []byte) (string, error) {
	data := excel.ConvertFromData(bytes.NewReader(fileData), int64(len(fileData)), "outline")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", base.NewEpError(1, fmt.Sprintf("Excel教材目录 json转换失败: %s", err))
	}

	return string(jsonBytes), nil
}


//读取 excel book 数据
func (c *CT_Excel) book(fileData []byte) (string, error) {
	data := excel.ConvertFromData(bytes.NewReader(fileData), int64(len(fileData)), "book")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", base.NewEpError(1, fmt.Sprintf("Excel书籍 json转换失败: %s", err))
	}

	return string(jsonBytes), nil
}

//读取 excel cognition_map 数据
func (c *CT_Excel) cognition_map(fileData []byte) (string, error) {
	data := excel.ConvertFromData(bytes.NewReader(fileData), int64(len(fileData)), "cognition_map")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", base.NewEpError(1, fmt.Sprintf("Excel普适认知点 json转换失败: %s", err))
	}

	return string(jsonBytes), nil
}

//读取 excel cognition_sp 数据
func (c *CT_Excel) cognition_sp(fileData []byte) (string, error) {
	data := excel.ConvertFromData(bytes.NewReader(fileData), int64(len(fileData)), "cognition_sp")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", base.NewEpError(1, fmt.Sprintf("Excel特异认知点 json转换失败: %s", err))
	}

	return string(jsonBytes), nil
}