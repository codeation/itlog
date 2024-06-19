#ifndef _SIGNALCONNECT_H_
#define _SIGNALCONNECT_H_

#include <gtk/gtk.h>

// Application
void app_startup(GApplication *application, gpointer data);
void app_activate(GApplication *application, gpointer data);
void app_shutdown(GApplication *application, gpointer data);

// Widget
gboolean window_draw(GtkWidget *widget, cairo_t *cr, gpointer data);
gboolean widget_delete(GtkWidget *widget, GdkEvent *event, gpointer data);
void widget_size_allocate(GtkWidget *widget, GtkAllocation *allocation, gpointer data);
gboolean widget_key_press(GtkWidget *widget, GdkEventKey *event, gpointer data);
gboolean widget_button_press(GtkWidget *widget, GdkEventButton *event, gpointer data);
gboolean widget_button_release(GtkWidget *widget, GdkEventButton *event, gpointer data);
gboolean widget_motion_notify(GtkWidget *widget, GdkEventMotion *event, gpointer data);
gboolean widget_scroll(GtkWidget *widget, GdkEventScroll *event, gpointer data);

// SimpleAction
void menu_item_activate(GSimpleAction *action, GVariant *parameter, gpointer data);

// Misc
void clipboard_text_received(GtkClipboard *clipboard, const gchar *text, gpointer data);
void object_weak_notify(gpointer data, GObject *where_the_object_was);

#endif
