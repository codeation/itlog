#ifndef _SIGNALCONNECT_H_
#define _SIGNALCONNECT_H_

#include <gtk/gtk.h>

// Application
void app_activate(GApplication *application, gpointer data);
void app_shutdown(GApplication *application, gpointer data);

// Widget
void window_draw(GtkDrawingArea *widget, cairo_t *cr, int width, int height, gpointer data);

gboolean close_request(GtkWindow *self, gpointer data);
void size_notify(GObject *self, GParamSpec *pspec, gpointer data);
void adjustment_notify(GtkAdjustment *self, gpointer data);
gboolean key_pressed(GtkEventControllerKey *self, guint keyval, guint keycode, GdkModifierType state,
                     gpointer data);
void button_pressed(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer data);
void button_released(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer data);
void motion_notify(GtkEventControllerMotion *self, gdouble x, gdouble y, gpointer data);
gboolean scroll_notify(GtkEventControllerScroll *self, gdouble dx, gdouble dy, gpointer data);

// SimpleAction
void menu_item_activate(GSimpleAction *action, GVariant *parameter, gpointer data);

// Clipboard
void clipboard_text_received(GObject *source_object, GAsyncResult *res, gpointer data);

// Misc
void object_weak_notify(gpointer data, GObject *where_the_object_was);

void size_notify_init();

#endif
