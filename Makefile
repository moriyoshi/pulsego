include $(GOROOT)/src/Make.$(GOARCH)

TARG=pulse
CGOFILES=\
	pulse.go
CGO_LDFLAGS=-lpulse
CLEANFILES+=test

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
