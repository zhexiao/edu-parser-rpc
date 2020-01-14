package word

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zhexiao/edu-parser-proto/basepb"
	"github.com/zhexiao/office-parser/word"
)

type CT_Word struct {
}

func (c *CT_Word) ParserPaper(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	data := word.ConvertPaperFromData(bytes.NewReader(fileBytes), int64(len(fileBytes)))

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		resp.Data = fmt.Sprintf("json转换失败: %s", err)
		return nil
	}

	resp.Data = string(jsonBytes)
	return nil
}

func (c *CT_Word) ParserQuestion(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	data := word.ConvertFromData(bytes.NewReader(fileBytes), int64(len(fileBytes)))

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		resp.Data = fmt.Sprintf("json转换失败: %s", err)
		return nil
	}

	resp.Data = string(jsonBytes)
	return nil
}
