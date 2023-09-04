package main

import (
	"context"
	"fmt"
	"log"

	"github.com/MohammadZeyaAhmad/Bank-App/api"
	db "github.com/MohammadZeyaAhmad/Bank-App/db/sqlc"
	"github.com/MohammadZeyaAhmad/Bank-App/mail"
	"github.com/MohammadZeyaAhmad/Bank-App/util"
	"github.com/MohammadZeyaAhmad/Bank-App/worker"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
)



func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
    runDBMigration(config.MigrationURL, config.DBSource)
	store := db.NewStore(connPool)
		redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
    
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(config, redisOpt, store)
	server, err := api.NewServer(config, store,taskDistributor)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.Port)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")

}

// func runGrpcServer(config util.Config, store db.Store) {
// 	server, err := gapi.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("cannot create server:", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterSimpleBankServer(grpcServer, server)
// 	reflection.Register(grpcServer)

// 	listener, err := net.Listen("tcp", config.GRPCServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot create listener:", err)
// 	}

// 	log.Printf("start gRPC server at %s", listener.Addr().String())
// 	err = grpcServer.Serve(listener)
// 	if err != nil {
// 		log.Fatal("cannot start gRPC server:", err)
// 	}
// }

// func runGatewayServer(config util.Config, store db.Store) {
// 	server, err := gapi.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("cannot create server:", err)
// 	}

// 	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
// 		MarshalOptions: protojson.MarshalOptions{
// 			UseProtoNames: true,
// 		},
// 		UnmarshalOptions: protojson.UnmarshalOptions{
// 			DiscardUnknown: true,
// 		},
// 	})

// 	grpcMux := runtime.NewServeMux(jsonOption)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err = pb.RegisterSimpleBankHandlerServer(ctx, grpcMux, server)
// 	if err != nil {
// 		log.Fatal("cannot register handler server:", err)
// 	}

// 	mux := http.NewServeMux()
// 	mux.Handle("/", grpcMux)

// 	listener, err := net.Listen("tcp", config.Port)
// 	if err != nil {
// 		log.Fatal("cannot create listener:", err)
// 	}

// 	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
// 	err = http.Serve(listener, mux)
// 	if err != nil {
// 		log.Fatal("cannot start HTTP gateway server:", err)
// 	}
// }

func runTaskProcessor(config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	fmt.Println("starting task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal("failed to start task processor")
	}
}

// func runGinServer(config util.Config, store db.Store) {
// 	// server, err := api.NewServer(config, store)
// 	if err != nil {
// 		log.Fatal("cannot create server:", err)
// 	}

// 	err = server.Start(config.Port)
// 	if err != nil {
// 		log.Fatal("cannot start server:", err)
// 	}
// }