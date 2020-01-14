package srv

import (
	"context"
	"github.com/zhexiao/edu-parser-rpc/proto/basepb"
)

func (c *CT_Word) ParserPaper(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	jsonData, err := c.paper(fileBytes)

	if err != nil {
		resp.Data = err.Error()
	} else {
		resp.Data = jsonData
	}

	return nil
}

func (c *CT_Word) ParserQuestion(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	jsonData, err := c.question(fileBytes)

	if err != nil {
		resp.Data = err.Error()
	} else {
		resp.Data = jsonData
	}

	return nil
}
