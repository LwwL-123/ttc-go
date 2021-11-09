package base

import "github.com/LwwL-123/go-substrate-rpc-client/v3/types"

type HrmpChannelId struct {
	Sender   types.U32
	Receiver types.U32
}
