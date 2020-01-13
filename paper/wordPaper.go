package paper

import (
	"context"
	"log"
)
import "github.com/zhexiao/edu-parser-proto/common"

type CT_WordPaper struct {
}

func (c *CT_WordPaper) Parser(ctx context.Context, req *common.Request, resp *common.Response) error {
	log.Println("word paper parser ...")

	resp.Data = "hello word"

	return nil
}
