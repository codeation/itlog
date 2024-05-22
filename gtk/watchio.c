#include "watchio.h"
#include <gtk/gtk.h>

void streamDo();

gboolean stream_read_chan(GIOChannel *source, GIOCondition condition, gpointer data) {
    streamDo();
    return TRUE;
}

void requestDo();

gboolean request_read_chan(GIOChannel *source, GIOCondition condition, gpointer data) {
    requestDo();
    return TRUE;
}

void chanErr();

gboolean chan_error_func(GIOChannel *source, GIOCondition condition, gpointer data) {
    chanErr();
    return TRUE;
}
