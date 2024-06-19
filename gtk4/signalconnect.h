#ifndef _SIGNALCONNECT_H_
#define _SIGNALCONNECT_H_

#include <gtk/gtk.h>

// Application
void app_startup(GApplication *application, gpointer data);
void app_activate(GApplication *application, gpointer data);
void app_shutdown(GApplication *application, gpointer data);

// Widget
void window_draw(GtkDrawingArea *widget, cairo_t *cr, int width, int height, gpointer data);

gboolean close_request(GtkWindow *self, gpointer user_data);
void size_notify(GObject *self, GParamSpec *pspec, gpointer user_data);
void adjustment_notify(GtkAdjustment *self, gpointer user_data);
gboolean key_pressed(GtkEventControllerKey *self, guint keyval, guint keycode, GdkModifierType state,
                     gpointer user_data);
void button_pressed(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer user_data);
void button_released(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer user_data);
void motion_notify(GtkEventControllerMotion *self, gdouble x, gdouble y, gpointer user_data);
gboolean scroll_notify(GtkEventControllerScroll *self, gdouble dx, gdouble dy, gpointer user_data);

// SimpleAction
void menu_item_activate(GSimpleAction *action, GVariant *parameter, gpointer data);

// Misc
void object_weak_notify(gpointer data, GObject *where_the_object_was);

void size_notify_init();

#endif
