package node

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	v1 "github.com/dun-io/imhub/api/imhub/node/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: v1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "QueryNodes",
					Use:       "queryNodes",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              v1.MsgService_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "MsgRegisterNode",
					Use:            "register-node [from] [address] [name] [pub_key]",
					Short:          "Send a register-node tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "from"}, {ProtoField: "address"}, {ProtoField: "name"}, {ProtoField: "pub_key"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
