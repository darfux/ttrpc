// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/ttrpc/example/example.proto
package example

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type ExampleService interface {
	Method1(ctx context.Context, req *Method1Request) (*Method1Response, error)
	Method2(ctx context.Context, req *Method1Request) (*empty.Empty, error)
}

func RegisterExampleService(srv *ttrpc.Server, svc ExampleService) {
	srv.Register("ttrpc.example.v1.Example", map[string]ttrpc.Method{
		"Method1": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req Method1Request
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.Method1(ctx, &req)
		},
		"Method2": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req Method1Request
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.Method2(ctx, &req)
		},
	})
}

type exampleClient struct {
	client *ttrpc.Client
}

func NewExampleClient(client *ttrpc.Client) ExampleService {
	return &exampleClient{
		client: client,
	}
}
func (c *exampleClient) Method1(ctx context.Context, req *Method1Request) (*Method1Response, error) {
	var resp Method1Response
	if err := c.client.Call(ctx, "ttrpc.example.v1.Example", "Method1", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (c *exampleClient) Method2(ctx context.Context, req *Method1Request) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "ttrpc.example.v1.Example", "Method2", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
