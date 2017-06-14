package rpcreceiver

import (
	"log"
	"net/rpc"

	"github.com/puppetlabs/lumogon/logging"
	"github.com/puppetlabs/lumogon/types"
)

// SendResult submits the ContainerReport populated with the output from each
// AttachedCapability run on the attached container.
// This sends to the `scheduler` host which will have been aliased at container
// creation to the node where the the scheduler is running.
func SendResult(result types.ContainerReport, harvesterHostname string) (bool, error) {
	client, err := rpc.Dial("tcp", "scheduler:42586")
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	var ack bool
	err = client.Call("RemoteMethods.SubmitCapabilities", result, &ack)
	if err != nil {
		logging.Debug("[RPC Client] Ack received: %t", ack)
		log.Fatal(err)
		return false, err
	}
	return true, nil
}
