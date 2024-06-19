#include "signalconnect.h"
#include <gtk/gtk.h>

/* Application */

void appStartup(GApplication *application, gpointer data);

void app_startup(GApplication *application, gpointer data) { appStartup(application, data); }

void appActivate(GApplication *application, gpointer data);

void app_activate(GApplication *application, gpointer data) { appActivate(application, data); }

void appShutdown(GApplication *application, gpointer data);

void app_shutdown(GApplication *application, gpointer data) { appShutdown(application, data); }

/* Widget */

void windowDraw(GtkWidget *widget, cairo_t *cr, gpointer data);

gboolean window_draw(GtkWidget *widget, cairo_t *cr, gpointer data) {
    windowDraw(widget, cr, data);
    return FALSE;
}

void widgetDelete(GtkWidget *widget, GdkEvent *event, gpointer data);

gboolean widget_delete(GtkWidget *widget, GdkEvent *event, gpointer data) {
    widgetDelete(widget, event, data);
    return TRUE;
}

void widgetSizeAllocate(GtkWidget *widget, GtkAllocation *allocation, gpointer data);

void widget_size_allocate(GtkWidget *widget, GtkAllocation *allocation, gpointer data) {
    widgetSizeAllocate(widget, allocation, data);
}

void widgetKeyPress(GtkWidget *widget, GdkEventKey *event, gpointer data);

gboolean widget_key_press(GtkWidget *widget, GdkEventKey *event, gpointer data) {
    widgetKeyPress(widget, event, data);
    return TRUE;
}

void widgetButtonPress(GtkWidget *widget, GdkEventButton *event, gpointer data);

gboolean widget_button_press(GtkWidget *widget, GdkEventButton *event, gpointer data) {
    widgetButtonPress(widget, event, data);
    return FALSE;
}

void widgetButtonRelease(GtkWidget *widget, GdkEventButton *event, gpointer data);

gboolean widget_button_release(GtkWidget *widget, GdkEventButton *event, gpointer data) {
    widgetButtonRelease(widget, event, data);
    return FALSE;
}

void widgetMotionNotify(GtkWidget *widget, GdkEventMotion *event, gpointer data);

gboolean widget_motion_notify(GtkWidget *widget, GdkEventMotion *event, gpointer data) {
    widgetMotionNotify(widget, event, data);
    return FALSE;
}

void widgetScroll(GtkWidget *widget, GdkEventScroll *event, gpointer data);

gboolean widget_scroll(GtkWidget *widget, GdkEventScroll *event, gpointer data) {
    widgetScroll(widget, event, data);
    return TRUE;
}

/* SimpleAction */

void menuItemActivate(GSimpleAction *action, GVariant *parameter, gpointer data);

void menu_item_activate(GSimpleAction *action, GVariant *parameter, gpointer data) {
    menuItemActivate(action, parameter, data);
}

/* Misc */

void clipboardTextReceived(GtkClipboard *clipboard, const gchar *text, gpointer data);

void clipboard_text_received(GtkClipboard *clipboard, const gchar *text, gpointer data) {
    clipboardTextReceived(clipboard, text, data);
}

void objectWeakRef(gpointer data, GObject *where_the_object_was);

void object_weak_notify(gpointer data, GObject *where_the_object_was) {
    // printf("object_weak_notify: %s %p\n", G_OBJECT_TYPE_NAME(where_the_object_was), where_the_object_was);
    objectWeakRef(data, where_the_object_was);
}
