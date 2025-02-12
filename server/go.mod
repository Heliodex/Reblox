module Reblox/server

go 1.22.3

replace Reblox/client v0.0.0 => ../client

replace Reblox/shared v0.0.0 => ../shared

require (
	Reblox/shared v0.0.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/ChainSafe/go-schnorrkel v1.1.0 // indirect
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d // indirect
	github.com/gtank/merlin v0.1.1-0.20191105220539-8318aed1a79f // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/mimoo/StrobeGo v0.0.0-20181016162300-f8f6d4d2b643 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
)
