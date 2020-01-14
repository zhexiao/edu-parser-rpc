package excel

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
