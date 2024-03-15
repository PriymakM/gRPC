package adder

import (
	"context"
	"fmt"
	"math"

	"grpc_pr3.1/pkg/proto/proto"
)

func factorial(n float64) float64 {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func count_data(x float64, a float64, b float64, c float64) float64 {

	size := 10
	var result float64

	for i := 0; i < size; i++ {
		result += (float64(math.Sin(x)) + float64(math.Cos(a-b))) / factorial((float64(i) + c))
	}

	return result
}

// GRPCServer......
type GRPCServer struct{}

// Add...
func (s *GRPCServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {

	var x, a, b, c float64

	x = req.GetX()
	fmt.Println(x)
	a = req.GetA()
	fmt.Println(a)
	b = req.GetB()
	fmt.Println(b)
	c = req.GetC()
	fmt.Println(c)

	result_equation := count_data(x, a, b, c)
	return &proto.AddResponse{Result: result_equation}, nil
}
