package hello

import (
	"log"
	"net/http"
	"os"

	"github.com/nsqio/go-nsq"
	"github.com/tokopedia/gosample/hello"
)

func ExampleCreateNewConsumer() {
	//creating nsq config
	nsqCfg := nsq.NewConfig()

	// calling CreateNewConsumer
	// handler is function which accept *nsq.Message as parameter
	q := CreateNewConsumer(nsqCfg, "random-topic", "test", handler)

	// set error level of the consumer
	q.SetLogger(log.New(os.Stderr, "nsq:", log.Ltime), nsq.LogLevelError)

	//connecting to nsq lookupd
	q.ConnectToNSQLookupd("nsqlookupd.local:4161")
}

func ExampleHelloWorldModule_SayHelloWorld() {
	//Initialize NewHelloWorldModule
	hwm := hello.NewHelloWorldModule()

	//use SayHelloWorld as http handler for /hello endpoint
	http.HandleFunc("/hello", hwm.SayHelloWorld)
}
