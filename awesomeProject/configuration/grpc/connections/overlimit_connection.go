package connections

import (
	"awesomeProject/adapter/output/protos/integrator"
	"google.golang.org/grpc"
	"log"
)

var OVERLIMIT_DIAL_URL integrator.SumClient

func OpenOverlimitDialConnection() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	OVERLIMIT_DIAL_URL = integrator.NewSumClient(conn)
}
