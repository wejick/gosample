package hello

import (
	"expvar"
	"log"
	"net/http"
	"os"

	"github.com/nsqio/go-nsq"
	logging "gopkg.in/tokopedia/logging.v1"
)

//ServerConfig hold server configuration
type ServerConfig struct {
	//Containts name information from [server] tag
	Name string
}

//Config main configuration container
type Config struct {
	//will hold information from [Server] tag inside config file
	Server ServerConfig
}

//HelloWorldModule main building blocks of this package
type HelloWorldModule struct {
	cfg       *Config
	q         *nsq.Consumer
	something string
	stats     *expvar.Int
}

//NewHelloWorldModule creates object of HelloWorldModule struct,
//it is not only creating the object but also :
//	* Reading configuration using logging package
//	* Creating new nsq consumer using CreateNewConsumer function
//	* Initiating stats as rspStats
func NewHelloWorldModule() *HelloWorldModule {

	var cfg Config

	ok := logging.ReadModuleConfig(&cfg, "config", "hello") || logging.ReadModuleConfig(&cfg, "files/etc/gosample", "hello")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

	// this message only shows up if app is run with -debug option, so its great for debugging
	logging.Debug.Println("hello init called", cfg.Server.Name)

	// example of creating nsq consumer
	nsqCfg := nsq.NewConfig()
	q := CreateNewConsumer(nsqCfg, "random-topic", "test", handler)
	q.SetLogger(log.New(os.Stderr, "nsq:", log.Ltime), nsq.LogLevelError)
	q.ConnectToNSQLookupd("nsqlookupd.local:4161")

	return &HelloWorldModule{
		cfg:       &cfg,
		something: "John Doe",
		stats:     expvar.NewInt("rpsStats"),
		q:         q,
	}

}

//SayHelloWorld is example of http handler, it accepts w http.ResponseWriter, r *http.Request as parameters
//In this function we increament the stats expvar counter
//
//How to use it
func (hlm *HelloWorldModule) SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	hlm.stats.Add(1)
	w.Write([]byte("Hello " + hlm.something))
}

func handler(msg *nsq.Message) error {
	log.Println("got message :", string(msg.Body))
	msg.Finish()
	return nil
}

//CreateNewConsumer creates new nsq cunsomer
//How to use this function is quite easy, we just need to provide it with
//nsqconfig, topic, channel name and function handler.
//
//Look at this example :
func CreateNewConsumer(nsqCfg *nsq.Config, topic string, channel string, handler nsq.HandlerFunc) *nsq.Consumer {
	q, err := nsq.NewConsumer(topic, channel, nsqCfg)
	if err != nil {
		log.Fatal("failed to create consumer for ", topic, channel, err)
	}
	q.AddHandler(handler)
	return q
}
