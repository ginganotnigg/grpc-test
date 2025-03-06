clean:
	rm proto/crypto/*.go

generate:
	buf generate --template buf.gen.yaml --path ./proto/crypto/crypto.proto