
// Code generated by Script. DO NOT EDIT.
// Source: script/codegen/main.go
//
// Generated by this command:
//
//	go run script/codegen/main.go

package zendesk

import "context"

func (z *Client) GetTicketsIterator(ctx context.Context, opts *PaginationOptions) *Iterator[Ticket] {
	return &Iterator[Ticket]{
		CommonOptions: opts.CommonOptions,
		pageSize:      opts.PageSize,
		hasMore:       true,
		isCBP:         opts.IsCBP,
		pageAfter:     "",
		pageIndex:     1,
		ctx:           ctx,
		obpFunc:       z.GetTicketsOBP,
		cbpFunc:       z.GetTicketsCBP,
	}
}

func (z *Client) GetTicketsOBP(ctx context.Context, opts *OBPOptions) ([]Ticket, Page, error) {
	var data struct {
		Tickets []Ticket `json:"tickets"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &OBPOptions{}
	}
	
	u, err := addOptions("/tickets.json", tmp)
	
	if err != nil {
		return nil, Page{}, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.Tickets, data.Page, nil
}

func (z *Client) GetTicketsCBP(ctx context.Context, opts *CBPOptions) ([]Ticket, CursorPaginationMeta, error) {
	var data struct {
		Tickets []Ticket `json:"tickets"`
		Meta    CursorPaginationMeta `json:"meta"`
	}

	tmp := opts
	if tmp == nil {
		tmp = &CBPOptions{}
	}
	
	u, err := addOptions("/tickets.json", tmp)
	
	if err != nil {
		return nil, data.Meta, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, data.Meta, err
	}
	return data.Tickets, data.Meta, nil
}
