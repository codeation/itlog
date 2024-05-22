CFLAGS=$(shell pkg-config --cflags gtk+-3.0)
LDFLAGS=$(shell pkg-config --libs gtk+-3.0)

format:
	clang-format -i gtk/macro.h
	clang-format -i gtk/signalconnect.c
	clang-format -i gtk/signalconnect.h
	clang-format -i gtk/watchio.c
	clang-format -i gtk/watchio.h
