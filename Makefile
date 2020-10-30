.EXPORT_ALL_VARIABLES:
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

NAME=callscraper
NUMVERSION=0.1.0
MINVERSION=`date -u +%Y%m%d%.H%M%S`

LDFLAGS=-ldflags "-s -w -X github.com/ruxton/amc_callscraper/versions.MINVERSION=${MINVERSION} -X github.com/ruxton/amc_callscraper/versions.VERSION=${VERSION}"
OUTDIR=pkg/$(VERSION)
OUTPATH=$(OUTDIR)/$(NAME).$(GOOS).$(GOARCH)

ifeq ($(GOOS),windows)
	OUTFILE=$(OUTPATH).exe
else
	OUTFILE=$(OUTPATH)
endif

all: build
build: VERSION=$(NUMVERSION)
run: VERSION=$(NUMVERSION)-DEV-$(MINVERSION)
deb: TMPPKGDIR=/tmp/tmpPackageDir
		 VERSION?=$(NUMVERSION)

ifeq ($(GOOS),linux)
build: main-build deb
else
build: main-build
endif

main-build:
	@go get -d
	@go build -a -v -trimpath $(LDFLAGS) -o $(OUTFILE) main.go cli.go
	@upx $(OUTFILE)

run:
	@go get
	@go run $(LDFLAGS) main.go cli.go

deb:
	@mkdir -p $(TMPPKGDIR)/DEBIAN
	@cp -R $(OUTDIR)/../../DEBIAN/* $(TMPPKGDIR)/DEBIAN/
	@trap 'rm -f tmpfile' 0 1 2 3 15
	@( echo 'cat <<END_OF_TEXT' && cat $(OUTDIR)/../../DEBIAN/control && echo 'END_OF_TEXT' ) > tmpfile
	@VERSION=${VERSION} source tmpfile > $(TMPPKGDIR)/DEBIAN/control
	@mkdir -p $(TMPPKGDIR)/usr/local/bin
	@cp $(OUTFILE) $(TMPPKGDIR)/usr/local/bin/$(NAME).linux
	@dpkg -b $(TMPPKGDIR) $(OUTPATH).deb
	@trap 'rm -r ${TMPPKGDIR}' 0 1 2 3 15
