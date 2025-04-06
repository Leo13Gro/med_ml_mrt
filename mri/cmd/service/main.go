// TODO: убрать мусор отсюда сделать нормальную инициализацию
package main

import (
	"context"
	"log/slog"
	"net"
	"os"

	"github.com/WantBeASleep/goooool/brokerlib"
	"github.com/WantBeASleep/goooool/grpclib"
	"github.com/WantBeASleep/goooool/loglib"

	pkgconfig "github.com/WantBeASleep/goooool/config"

	"mri/internal/config"

	"mri/internal/repository"

	devicesrv "mri/internal/services/device"
	imagesrv "mri/internal/services/image"
	mrisrv "mri/internal/services/mri"
	nodesrv "mri/internal/services/node"
	segmentsrv "mri/internal/services/segment"

	pb "mri/internal/generated/grpc/service"
	grpchandler "mri/internal/grpc"

	devicehandler "mri/internal/grpc/device"
	imagehandler "mri/internal/grpc/image"
	mrihandler "mri/internal/grpc/mri"
	nodehandler "mri/internal/grpc/node"
	segmenthandler "mri/internal/grpc/segment"

	mriprocessedsubscriber "mri/internal/subs/mriprocessed"
	mriuploadsubscriber "mri/internal/subs/mriupload"

	adapters "mri/internal/adapters"
	brokeradapter "mri/internal/adapters/broker"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"
)

const (
	successExitCode = 0
	failExitCode    = 1
)

func main() {
	os.Exit(run())
}

func run() (exitCode int) {
	loglib.InitLogger(loglib.WithDevEnv())
	cfg, err := pkgconfig.Load[config.Config]()
	if err != nil {
		slog.Error("init config", "err", err)
		return failExitCode
	}

	db, err := sqlx.Open("postgres", cfg.DB.Dsn)
	if err != nil {
		slog.Error("init db", "err", err)
		return failExitCode
	}
	defer db.Close()

	client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Secure: false,
		Creds:  credentials.NewStaticV4(cfg.S3.Access_Token, cfg.S3.Secret_Token, ""),
	})
	if err != nil {
		slog.Error("init s3", "err", err)
		return failExitCode
	}

	if err := db.Ping(); err != nil {
		slog.Error("ping db", "err", err)
		return failExitCode
	}

	producer, err := brokerlib.NewProducer(cfg.Broker.Addrs)
	if err != nil {
		slog.Error("init broker producer", slog.Any("err", err))
	}

	dao := repository.NewRepository(db, client, "mri")
	adapter := adapters.New(brokeradapter.New(producer))

	deviceSrv := devicesrv.New(dao)
	mriSrv := mrisrv.New(dao)
	imageSrv := imagesrv.New(dao, adapter)
	nodeSrv := nodesrv.New(dao)
	serviceSrv := segmentsrv.New(dao)

	// grpc
	deviceHandler := devicehandler.New(deviceSrv)
	mriHandler := mrihandler.New(mriSrv)
	imageHandler := imagehandler.New(imageSrv)
	nodeHandler := nodehandler.New(nodeSrv)
	serviceHandler := segmenthandler.New(serviceSrv)

	handler := grpchandler.New(
		deviceHandler,
		mriHandler,
		imageHandler,
		nodeHandler,
		serviceHandler,
	)

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpclib.ServerCallPanicRecoverInterceptor,
			grpclib.ServerCallLoggerInterceptor,
		),
	)
	pb.RegisterMriSrvServer(server, handler)

	// broker
	mriuploadSubscriber := mriuploadsubscriber.New(imageSrv)
	mriprocessedSubscriber := mriprocessedsubscriber.New(nodeSrv)

	mriuploadHandler, err := brokerlib.GetSubscriberHandler(
		mriuploadSubscriber,
		cfg.Broker.Addrs,
		nil,
	)
	if err != nil {
		slog.Error("create mripload sub", "err", err)
		return failExitCode
	}

	mriprocessedHandler, err := brokerlib.GetSubscriberHandler(
		mriprocessedSubscriber,
		cfg.Broker.Addrs,
		nil,
	)
	if err != nil {
		slog.Error("create mriprocesse sub", "err", err)
		return failExitCode
	}

	lis, err := net.Listen("tcp", cfg.App.Url)
	if err != nil {
		slog.Error("take port", "err", err)
		return failExitCode
	}

	close := make(chan struct{})
	// ЛЮТОЕ MVP
	slog.Info("start serve", slog.String("app url", cfg.App.Url))
	go func() {
		if err := server.Serve(lis); err != nil {
			slog.Error("take port", "err", err)
			panic("serve grpc")
		}
		close <- struct{}{}
	}()
	go func() {
		// пока без DI
		if err := mriuploadHandler.Start(context.Background()); err != nil {
			slog.Error("start mriupload handler", "err", err)
		}
	}()
	go func() {
		if err := mriprocessedHandler.Start(context.Background()); err != nil {
			slog.Error("start mriprocessedHandler handler", "err", err)
		}
	}()

	<-close

	return successExitCode
}
