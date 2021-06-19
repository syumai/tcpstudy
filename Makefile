test/tcp/server:
	go test ./example/tcp/server/... -count=1

test/tcp/client:
	go test ./example/tcp/client/... -count=1
