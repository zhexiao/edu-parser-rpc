package word

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zhexiao/edu-parser-rpc/base"
	"github.com/zhexiao/office-parser/word"
)

type CT_Word struct {
}

func NewCT_Word() *CT_Word {
	return &CT_Word{}
}

//读取 word paper 数据
func (c *CT_Word) paper(fileData []byte) (string, error) {
	data := word.ConvertPaperFromData(bytes.NewReader(fileData), int64(len(fileData)))
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", base.NewEpError(1, fmt.Sprintf("Word试卷 json转换失败: %s", err))
	}

	return string(jsonBytes), nil
}

//读取 word question 数据
func (c *CT_Word) question(fileData []byte) (string, error) {
	data := word.ConvertFromData(bytes.NewReader(fileData), int64(len(fileData)))
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", base.NewEpError(1, fmt.Sprintf("Word试题 json转换失败: %s", err))
	}

	return string(jsonBytes), nil
}
