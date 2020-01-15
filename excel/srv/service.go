package srv

import (
	"context"
	"github.com/zhexiao/edu-parser-rpc/proto/basepb"
)

func (c *CT_Excel) ParserPaper(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	jsonData, err := c.paper(fileBytes)

	if err != nil {
		resp.Data = err.Error()
	} else {
		resp.Data = jsonData
	}

	return nil
}

func (c *CT_Excel) ParserOutline(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	jsonData, err := c.outline(fileBytes)

	if err != nil {
		resp.Data = err.Error()
	} else {
		resp.Data = jsonData
	}

	return nil
}

func (c *CT_Excel) ParserBook(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	jsonData, err := c.book(fileBytes)

	if err != nil {
		resp.Data = err.Error()
	} else {
		resp.Data = jsonData
	}

	return nil
}

func (c *CT_Excel) ParserCognitionSp(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	jsonData, err := c.cognition_sp(fileBytes)

	if err != nil {
		resp.Data = err.Error()
	} else {
		resp.Data = jsonData
	}

	return nil
}

func (c *CT_Excel) ParserCognitionMap(ctx context.Context, req *basepb.Request, resp *basepb.Response) error {
	fileBytes := req.Body
	jsonData, err := c.cognition_map(fileBytes)

	if err != nil {
		resp.Data = err.Error()
	} else {
		resp.Data = jsonData
	}

	return nil
}
