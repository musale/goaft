package tests

import "testing"
import "github.com/musale/goaft/lib"
import "fmt"

// TestHey test Hey()
func TestSendingMessage(t *testing.T) {
	gateway := lib.AfricastalkingGateway{Username: "test", APIKEY: "test", Debug: true, Format: "json"}
	response, err := gateway.SendMessage("0705867162", "message", "TEST")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	fmt.Println(response)
}
