package main

import (
	"net/http"

	"github.com/gorilla/rpc"
	jsonrpc "github.com/gorilla/rpc/json"
	jsonrpc "github.com/gorilla/rpc/v2/json"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("go-filecoin")

func init() {
	logging.SetAllLoggers(logging.LevelInfo)
}

type ConfigReply struct {
	TrackGas         bool `json:"trackGas"`
	CheckExitCode    bool `json:"checkExitCode"`
	CheckReturnValue bool `json:"checkReturnValue"`
	CheckStateRoot   bool `json:"checkStateRoot"`

	TestSuite []string `json:"testSuite"`
}

type ConfigRequest struct {
	Method string `json:"method"`
	Id     string `json:"id"`
}

type ConfigResponse struct {
	Result ConfigReply `json:"result"`
	Error  string      `json:"error"`
	Id     string      `json:"id"`
}

type ConfigService struct{}

func (c *ConfigService) Config(r *http.Request, args *struct{}, reply *ConfigReply) error {
	log.Infow("got requrest", "request", r.Host)
	reply.CheckExitCode = true
	reply.CheckReturnValue = true
	reply.CheckStateRoot = true
	reply.TrackGas = false
	reply.TestSuite = []string{"TestOne", "TestTwo", "TestThree"}
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	if err := s.RegisterService(new(ConfigService), ""); err != nil {
		panic(err)
	}
	http.Handle("/rpc", s)
	if err := http.ListenAndServe("127.0.0.1:8378", nil); err != nil {
		panic(err)
	}
}
