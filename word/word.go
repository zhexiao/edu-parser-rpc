package word

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/zhexiao/edu-parser-proto/basepb"
	"github.com/zhexiao/office-parser/word"
	"log"
)

type CT_Word struct {
}

func (c *CT_Word) ParserPaper(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	log.Println("word paper parser ...")

	fileBytes := req.Body
	data := word.ConvertPaperFromData(bytes.NewReader(fileBytes), int64(len(fileBytes)))

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Panicf("json转换失败: %s", err)
	}
	resp.Data = string(jsonBytes)

	return nil
}

func (c *CT_Word) ParserQuestion(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	log.Println("word question parser ...")
	fileBytes := req.Body

	data := word.ConvertFromData(bytes.NewReader(fileBytes), int64(len(fileBytes)))

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Panicf("json转换失败: %s", err)
	}
	resp.Data = string(jsonBytes)

	return nil
}
