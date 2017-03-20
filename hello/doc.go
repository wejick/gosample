// Documentation Tips : Any comment before package will be put as overview in your documentation

/*
Package hello can be used as reference when creating new module.

This package containts implementation of http handler, consumer and logging usage.
The main part of HelloWorldModule struct which is initiated by NewHelloWorldModule() method.
Usage :
	hwm := hello.NewHelloWorldModule()
	http.HandleFunc("/hello", hwm.SayHelloWorld)

By calling NewHelloWorldModule() we are creating object of HelloWorldModule, which allow us
to use SayHelloWorld as http handler. This is part of basic object oriented style of golang,
find more about it here : https://code.tutsplus.com/tutorials/lets-go-object-oriented-programming-in-golang--cms-26540.

Inside NewHelloWorldModule method configuration is loaded using logging package :

	var cfg Config

	ok := logging.ReadModuleConfig(&cfg, "config", "hello") || logging.ReadModuleConfig(&cfg, "files/etc/gosample", "hello")
	if !ok {
		// when the app is run with -e switch, this message will automatically be redirected to the log file specified
		log.Fatalln("failed to read config")
	}

Here logging will read config file named hello.[env].ini in files/etc/gosample and put in inside cfg.
Logging will read TKPENV environtment variable and when it's not set, the default value is development.
*/
package hello
