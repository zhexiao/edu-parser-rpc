package paper

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/zhexiao/office-parser/word"
	"log"
)
import "github.com/zhexiao/edu-parser-proto/common"

type CT_WordPaper struct {
}

func (c *CT_WordPaper) Parser(ctx context.Context, req *common.Request, resp *common.Response) error {
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
