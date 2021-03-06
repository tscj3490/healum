package main

import (
	"context"
	account_proto "server/account-srv/proto/account"
	"server/common"
	"server/content-srv/db"
	"server/content-srv/handler"
	content_proto "server/content-srv/proto/content"
	kv_proto "server/kv-srv/proto/kv"
	pubsub_proto "server/static-srv/proto/pubsub"
	static_proto "server/static-srv/proto/static"
	team_proto "server/team-srv/proto/team"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-os/config"
	"github.com/micro/go-os/config/source/file"
	"github.com/micro/go-os/metrics"
	_ "github.com/micro/go-plugins/broker/nats"
	"github.com/micro/go-plugins/metrics/telegraf"
	_ "github.com/micro/go-plugins/transport/nats"
	log "github.com/sirupsen/logrus"
)

func main() {
	configFile, _ := common.PopParameter("config")
	// Create a config instance
	conf := config.NewConfig(
		// poll every hour
		config.PollInterval(time.Hour),
		// use file as a config source
		config.WithSource(file.NewSource(config.SourceName(configFile))),
	)

	defer conf.Close()

	// create new metrics
	m := telegraf.NewMetrics(
		metrics.Namespace(conf.Get("service", "name").String("content")),
		metrics.Collectors(
			// telegraf/statsd address
			common.MetricAddress(),
		),
	)
	defer m.Close()

	common.ContentSrv = conf.Get("service", "content").String("go.micro.srv.content")
	version := conf.Get("service", "version").String("latest")
	descr := conf.Get("service", "description").String("Micro service")
	service := micro.NewService(
		micro.Name(common.ContentSrv),
		micro.Version(version),
		micro.Metadata(map[string]string{"Description": descr}),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*10),
		micro.Flags(
			cli.BoolFlag{
				Name: "debug",
			},
		),
	)
	var debugEnabled bool
	service.Init(
		micro.Action(func(c *cli.Context) {
			debugEnabled = c.Bool("debug")
		}),
	)
	if debugEnabled {
		log.SetLevel(log.DebugLevel)
	}

	brker := service.Client().Options().Broker
	brker.Connect()
	contentService := &handler.ContentService{
		Broker:        brker,
		StaticClient:  static_proto.NewStaticServiceClient("go.micro.srv.static", service.Client()),
		AccountClient: account_proto.NewAccountServiceClient("go.micro.srv.account", service.Client()),
		KvClient:      kv_proto.NewKvServiceClient("go.micro.srv.kv", service.Client()),
		TeamClient:    team_proto.NewTeamServiceClient("go.micro.srv.team", service.Client()),
	}
	content_proto.RegisterContentServiceHandler(service.Server(), contentService)

	// subscribe in the channels
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_ACTIVITY_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_RECIPE_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_EXCERCISE_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_ARTICLE_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_PLACE_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_WELLBEING_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_CONTENT_RECOMMENDATION}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_VIDEO_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_PRODUCT_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_SERVICE_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_EVENT_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}
	if err := contentService.Subscribe(context.TODO(), &pubsub_proto.SubscribeRequest{common.CREATE_RESEARCH_CONTENT}, &pubsub_proto.SubscribeResponse{}); err != nil {
		log.Fatal(err)
	}

	go contentService.WarmupCacheContent(context.TODO(), &content_proto.WarmupCacheContentRequest{}, &content_proto.WarmupCacheContentResponse{})

	db.Init(service.Client())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
