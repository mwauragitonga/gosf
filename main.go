package main

import (
  "fmt"
  "log"
  gosf "github.com/ambelovsky/gosf"

)

func init() {
  // Listen on an endpoint
  gosf.Listen("echo", func(client *gosf.Client, request *gosf.Request) *gosf.Message {
  
    return gosf.NewSuccessMessage(request.Message.Text)
  })
  gosf.Listen("sample", myFirstController)

  //listen to new socket connection event from client
    gosf.Listen("newConn", newConnController)

  gosf.OnConnect(func(client *gosf.Client, request *gosf.Request) {
  // Do something when a client connects
   

})

}

func main() {
  // Start the server using a basic configuration
  gosf.Startup(map[string]interface{}{
    "port": 9999})
    fmt.Println("server running on port 9999")

}

func myFirstController(client *gosf.Client, request *gosf.Request) *gosf.Message {
  var statement string

  // Parsing arguments in the body element
  if val, ok := request.Message.Body["statement"]; ok {
    statement = val.(string)
  }

  // Default value
  if statement == "" {
    statement = "that I love you."
  }
  
  // Construct response
  response := new(gosf.Message)
  response.Success = true
  response.Text = "Hello World, I'd just like to say " + statement

  // Send response back to the client
  return response
  
}
func newConnController(client *gosf.Client, request *gosf.Request) *gosf.Message {
  log.Printf("New Connection Details received:")
  var socket_id string

  // Parsing arguments in the body element
  if val, ok := request.Message.Body["socketId"]; ok {
    socket_id = val.(string)
  }
 var text = request.Message.Text
 fmt.Println(text)
  //print out the id
  log.Printf("Latest Newly Connected Socket ID is ---- %v",socket_id)
  //return response
  return gosf.NewSuccessMessage("Socket ID received By Server")
}
