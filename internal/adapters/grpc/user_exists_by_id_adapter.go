package grpc

import (
	"context"
	"log"
	"time"

	"gitlab.com/velo-company/services/events-service/internal/core/ports"
	"gitlab.com/velo-company/services/events-service/proto/user"
	"google.golang.org/grpc"
)

type userExistsByIdAdapter struct {
	client user.UserServiceClient
}

func NewUserExistsByIdAdapter(connection *grpc.ClientConn) ports.UserExistsByIdPort {
	return &userExistsByIdAdapter{
		client: user.NewUserServiceClient(connection),
	}
}

func (u userExistsByIdAdapter) Execute(userId int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.client.UserExistsById(ctx, &user.UserExistsByIdRequest{
		Id: int32(userId),
	})

	if err != nil {
		log.Print(err.Error())
		return false, err
	}

	return res.Exists, nil
}
