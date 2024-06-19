#ifndef _WATCHIO_H_
#define _WATCHIO_H_

#include <gtk/gtk.h>

gboolean stream_read_chan(GIOChannel *source, GIOCondition condition, gpointer data);
gboolean request_read_chan(GIOChannel *source, GIOCondition condition, gpointer data);
gboolean chan_error_func(GIOChannel *source, GIOCondition condition, gpointer data);

#endif
