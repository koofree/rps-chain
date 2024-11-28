package rps

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	rpsv1 "github.com/0xlb/rpschain/api/rps/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: nil,
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: rpsv1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CreateGame",
					Use:       "create opponent rounds",
					Short:     "Creates a new Rock, Paper & Scissors game for the message sender and the chose opponent.",
					Long:      "Creates a new Rock, Paper & Scissors game for the message sender and the chose opponent. Input parameters are the opponent address and the rounds",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "opponent"},
						{ProtoField: "rounds"},
					},
				},
			},
		},
	}
}
