clean:
	rm proto/hello/*.go

generate:
	buf generate --path ./proto/hello/hello.proto