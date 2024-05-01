
VERSION := v0.1.0

BINDIR := bin


ALL_BIN := \
	$(BINDIR)/gostretch-$(VERSION)-linux-amd64 \
	$(BINDIR)/gostretch-$(VERSION)-linux-arm64 \
	$(BINDIR)/gostretch-$(VERSION)-darwin-amd64 \
	$(BINDIR)/gostretch-$(VERSION)-darwin-arm64 \
	$(BINDIR)/gostretch-$(VERSION)-windows-amd64 \
	$(BINDIR)/gostretch-$(VERSION)-windows-arm64

.PHONY: all $(ALL_BIN)

all: $(BINDIR) $(ALL_BIN)

bin:
	mkdir -p bin

$(BINDIR)/gostretch-$(VERSION)-linux-amd64: cmd/main.go
	GOOS=linux GOARCH=amd64 go build -o $@ $<

$(BINDIR)/gostretch-$(VERSION)-linux-arm64: cmd/main.go
	GOOS=linux GOARCH=arm64 go build -o $@ $<

$(BINDIR)/gostretch-$(VERSION)-darwin-amd64: cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o $@ $<

$(BINDIR)/gostretch-$(VERSION)-darwin-arm64: cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o $@ $<

$(BINDIR)/gostretch-$(VERSION)-windows-amd64: cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o $@ $<

$(BINDIR)/gostretch-$(VERSION)-windows-arm64: cmd/main.go
	GOOS=windows GOARCH=arm64 go build -o $@ $<
