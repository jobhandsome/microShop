// Code generated by goctl. DO NOT EDIT!
// Source: cart.proto

package server

import (
	"context"

	"microShop/cart/rpc/internal/logic"
	"microShop/cart/rpc/internal/svc"
	"microShop/cart/rpc/types/cart"
)

type CartServer struct {
	svcCtx *svc.ServiceContext
	cart.UnimplementedCartServer
}

func NewCartServer(svcCtx *svc.ServiceContext) *CartServer {
	return &CartServer{
		svcCtx: svcCtx,
	}
}

func (s *CartServer) AddCart(ctx context.Context, in *cart.AddCartReq) (*cart.CommonResply, error) {
	l := logic.NewAddCartLogic(ctx, s.svcCtx)
	return l.AddCart(in)
}

func (s *CartServer) EditCart(ctx context.Context, in *cart.EditCartReq) (*cart.CommonResply, error) {
	l := logic.NewEditCartLogic(ctx, s.svcCtx)
	return l.EditCart(in)
}

func (s *CartServer) DelCart(ctx context.Context, in *cart.DelCartReq) (*cart.CommonResply, error) {
	l := logic.NewDelCartLogic(ctx, s.svcCtx)
	return l.DelCart(in)
}
