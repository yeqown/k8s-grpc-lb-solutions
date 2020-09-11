build-sniff:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sniff ./sniff_01.go
	docker build -t "yeqown/sniff:v0.0.1" . -f sniff.Dockerfile
	docker push "yeqown/sniff:v0.0.1"
