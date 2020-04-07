package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/filecoin-project/chain-validation/chain/types"
	"github.com/filecoin-project/chain-validation/state"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/filecoin-project/specs-actors/actors/builtin/power"
	rpc "github.com/gorilla/rpc/v2"
	jsonrpc "github.com/gorilla/rpc/v2/json"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log"
	cbg "github.com/whyrusleeping/cbor-gen"
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

type ConfigService struct{}

func (c *ConfigService) Config(r *http.Request, args *struct{}, reply *ConfigReply) error {
	reply.CheckExitCode = true
	reply.CheckReturnValue = true
	reply.CheckStateRoot = true
	reply.TrackGas = false
	reply.TestSuite = []string{"TestOne", "TestTwo", "TestThree"}

	log.Infow("ConfigService.Config", "reply", reply)
	return nil
}

type ApplyMessageArgs struct {
	Epoch   abi.ChainEpoch
	Message *types.Message
}

type ApplySignedMessageArgs struct {
	Epoch         abi.ChainEpoch
	SignedMessage *types.SignedMessage
}

type ApplyMessageReply struct {
	Receipt types.MessageReceipt
	Penalty abi.TokenAmount
	Reward  abi.TokenAmount
}

type ApplyTipSetMessagesArgs struct {
	Epoch       abi.ChainEpoch
	ParentState cid.Cid
	Blocks      []types.BlockMessagesInfo
	Randomness  abi.Randomness
}

type ApplyTipSetMessagesReply struct {
	Receipts []types.MessageReceipt
}

type ApplierService struct {
	applier state.Applier
}

func (a *ApplierService) ApplyMessage(r *http.Request, args *ApplyMessageArgs, reply *ApplyMessageReply) error {
	log.Infow("ApplierService.ApplyMessage", "args", args, "reply", reply)
	return nil
}

func (a *ApplierService) ApplySignedMessage(r *http.Request, args *ApplySignedMessageArgs, reply *ApplyMessageReply) error {
	log.Infow("ApplierService.ApplySignedMessage", "args", args, "reply", reply)
	return nil
}

func (a *ApplierService) ApplyTipSetMessages(r *http.Request, args *ApplyTipSetMessagesArgs, reply *ApplyTipSetMessagesReply) error {
	log.Infow("ApplierService.ApplyTipSetMessages", "args", args, "reply", reply)
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(jsonrpc.NewCodec(), "application/json")
	if err := s.RegisterService(new(ConfigService), ""); err != nil {
		panic(err)
	}
	if err := s.RegisterService(new(ApplierService), ""); err != nil {
		panic(err)
	}
	http.Handle("/rpc", s)
	if err := http.ListenAndServe("127.0.0.1:8378", nil); err != nil {
		panic(err)
	}
}
