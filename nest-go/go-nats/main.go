package main 

import (
	"github.com/joho/godotenv"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn
var ec *nats.EncodedConn


func init(){
	nc, err = nats.Connect(os.Getenv(("NATS_URI")))
	ec, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	// creating subscription 
	nc.Subscribe("foo", func(msg *nats.Msg) {
		msg.Respond([]byte("OK OK OK"))
	})

	// publish to an event in nest service 
	// psudo_code 
	var data map[string]interface{}{} // jsondata 
	ec.publish("create-data-nats",&data) 
}

func main(){


	
}