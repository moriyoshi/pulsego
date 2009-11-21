include $(GOROOT)/src/Make.$(GOARCH)

TARG=pulse
CGOFILES=\
	pulse.go
CGO_LDFLAGS=-lpulse
CLEANFILES+=hello

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
