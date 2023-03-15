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

  //listen to connection event from client
    gosf.Listen("newConn", newConnController)

  gosf.OnConnect(func(client *gosf.Client, request *gosf.Request) {
  // Do something when a client connects
   

})
  gosf.RegisterPlugin(new(Plugin))
	log.Println("Sample Plugin Initialized")
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
  log.Printf("New Connection Details:")
  var statement string

  // Parsing arguments in the body element
  if val, ok := request.Message.Body["id"]; ok {
    statement = val.(string)
  }
 var text = request.Message.Text
 fmt.Println(text)
  //print out the id
  log.Printf("here is the %v",statement)
  //return response
  return gosf.NewSuccessMessage("New Connection Details Sent To Server")
}
/** ASPECT **/

// Plugin is the aspect oriented element required by the modular plugin framework
type Plugin struct{}

// Activate is an aspect-oriented modular plugin requirement
func (p Plugin) Activate(app *gosf.AppSettings) {}

// Deactivate is an aspect-oriented modular plugin requirement
func (p Plugin) Deactivate(app *gosf.AppSettings) {}