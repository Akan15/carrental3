package test

import (
	"context"
	"testing"

	rental "github.com/Akan15/carrental3/rental-service/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestRentalServiceIntegration(t *testing.T) {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	assert.NoError(t, err)
	defer conn.Close()

	client := rental.NewRentalServiceClient(conn)

	createResp, err := client.CreateRental(context.Background(), &rental.CreateRentalRequest{
		UserId: "user1", CarId: "car1", Type: "normal",
	})
	assert.NoError(t, err)
	assert.Equal(t, "user1", createResp.UserId)

	endResp, err := client.EndRental(context.Background(), &rental.EndRentalRequest{Id: createResp.Id})
	assert.NoError(t, err)
	assert.NotEmpty(t, endResp.EndTime)

	getResp, err := client.GetRental(context.Background(), &rental.GetRentalRequest{Id: createResp.Id})
	assert.NoError(t, err)
	assert.Equal(t, createResp.Id, getResp.Id)
}
