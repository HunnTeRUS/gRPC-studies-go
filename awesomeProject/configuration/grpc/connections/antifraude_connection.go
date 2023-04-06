package connections

import (
	"awesomeProject/adapter/output/protos/integrator"
	"google.golang.org/grpc"
	"log"
)

var ANTIFRAUD_DIAL_URL integrator.SumClient

func OpenAntifraudDialConnection() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	ANTIFRAUD_DIAL_URL = integrator.NewSumClient(conn)
}
