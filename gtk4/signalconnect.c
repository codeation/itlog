#include "signalconnect.h"
#include <gtk/gtk.h>

/* Application */

void appActivate(GApplication *application, gpointer data);

void app_activate(GApplication *application, gpointer data) { appActivate(application, data); }

void appShutdown(GApplication *application, gpointer data);

void app_shutdown(GApplication *application, gpointer data) { appShutdown(application, data); }

/* Widget */

void windowDraw(GtkDrawingArea *widget, cairo_t *cr, gpointer data);

void window_draw(GtkDrawingArea *widget, cairo_t *cr, int width, int height, gpointer data) {
    windowDraw(widget, cr, data);
}

void widgetDelete(GtkWindow *self, gpointer data);

gboolean close_request(GtkWindow *self, gpointer data) {
    widgetDelete(self, data);
    return TRUE;
}

int idle_count = 0;

void widgetIdle(gpointer data);

gboolean idle_func(gpointer data) {
    idle_count--;
    if (idle_count > 0) {
        return G_SOURCE_CONTINUE;
    }
    widgetIdle(data);
    return G_SOURCE_REMOVE;
}

void size_notify(GObject *self, GParamSpec *pspec, gpointer data) {
    if (idle_count <= 0) {
        g_idle_add(idle_func, data);
    }
    idle_count++;
}

void adjustment_notify(GtkAdjustment *self, gpointer data) {
    if (idle_count <= 0) {
        g_idle_add(idle_func, data);
    }
    idle_count++;
}

void size_notify_init(gpointer data) { adjustment_notify(NULL, data); }

void widgetKeyPress(GtkEventControllerKey *self, guint keyval, guint keycode, GdkModifierType state,
                    gpointer data);

gboolean key_pressed(GtkEventControllerKey *self, guint keyval, guint keycode, GdkModifierType state,
                     gpointer data) {
    widgetKeyPress(self, keyval, keycode, state, data);
    return TRUE;
}

void widgetButtonPress(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer data);

void button_pressed(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer data) {
    widgetButtonPress(self, n_press, x, y, data);
}

void widgetButtonRelease(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer data);

void button_released(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer data) {
    widgetButtonRelease(self, n_press, x, y, data);
}

void widgetMotionNotify(GtkEventControllerMotion *self, gdouble x, gdouble y, gpointer data);

void motion_notify(GtkEventControllerMotion *self, gdouble x, gdouble y, gpointer data) {
    widgetMotionNotify(self, x, y, data);
}

void widgetScroll(GtkEventControllerScroll *self, gdouble dx, gdouble dy, gpointer data);

gboolean scroll_notify(GtkEventControllerScroll *self, gdouble dx, gdouble dy, gpointer data) {
    widgetScroll(self, dx, dy, data);
    return TRUE;
}

/* SimpleAction */

void menuItemActivate(GSimpleAction *action, GVariant *parameter, gpointer data);

void menu_item_activate(GSimpleAction *action, GVariant *parameter, gpointer data) {
    menuItemActivate(action, parameter, data);
}

/* Clipboard */

void clipboardTextReceived(GObject *source_object, GAsyncResult *res, gpointer data);

void clipboard_text_received(GObject *source_object, GAsyncResult *res, gpointer data) {
    clipboardTextReceived(source_object, res, data);
}

/* Misc */

void objectWeakRef(gpointer data, GObject *where_the_object_was);

void object_weak_notify(gpointer data, GObject *where_the_object_was) {
    // printf("object_weak_notify: %s %p\n", G_OBJECT_TYPE_NAME(where_the_object_was), where_the_object_was);
    objectWeakRef(data, where_the_object_was);
}

void idleNotify();

gboolean idle_notify(gpointer data) {
    idleNotify();
    return G_SOURCE_CONTINUE;
}
