package handlers

import (
	"context"

	pb "github.com/Serjeri/proto-exchange/exchange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	pb.UnimplementedExchangeServiceServer
	user UserService
}

type UserService interface {
	GetExchange(ctx context.Context) (map[string]float64, error)
	GetRate(ctx context.Context, fromCurrency, toCurrency string, amount int) (float64, error)
}

func Register(gRPCServer *grpc.Server, user UserService) {
	pb.RegisterExchangeServiceServer(gRPCServer, &serverAPI{user: user})
}

func (s *serverAPI) GetExchangeRates(ctx context.Context, in *pb.Empty) (*pb.ExchangeRatesResponse, error) {
	rates, err := s.user.GetExchange(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed get exchange rates")
	}

	return &pb.ExchangeRatesResponse{Rates: rates}, nil
}

func (s *serverAPI) PerformExchange(ctx context.Context, in *pb.ExchangeRequest) (*pb.ExchangeResponse, error) {
	rate, err := s.user.GetRate(ctx, in.FromCurrency, in.ToCurrency, int(in.Amount))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed get perform exchange ")
	}

	newBalance := make(map[string]float64)
	newBalance[in.FromCurrency] = 0.00
	newBalance[in.ToCurrency] = rate

	return &pb.ExchangeResponse{
		Message:         "Exchange successful",
		ExchangedAmount: float32(rate),
		NewBalance:      newBalance,
	}, nil
}
