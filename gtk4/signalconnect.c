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

void windowDraw(GtkDrawingArea *widget, cairo_t *cr, gpointer data);

void window_draw(GtkDrawingArea *widget, cairo_t *cr, int width, int height, gpointer data) {
    windowDraw(widget, cr, data);
}

void widgetDelete(GtkWindow *self, gpointer user_data);

gboolean close_request(GtkWindow *self, gpointer user_data) {
    widgetDelete(self, user_data);
    return TRUE;
}

int idle_count = 0;

void widgetIdle(gpointer user_data);

gboolean idle_func(gpointer user_data) {
    idle_count--;
    if (idle_count > 0) {
        return G_SOURCE_CONTINUE;
    }
    widgetIdle(user_data);
    return G_SOURCE_REMOVE;
}

void size_notify(GObject *self, GParamSpec *pspec, gpointer user_data) {
    if (idle_count <= 0) {
        g_idle_add(idle_func, user_data);
    }
    idle_count++;
}

void adjustment_notify(GtkAdjustment *self, gpointer user_data) {
    if (idle_count <= 0) {
        g_idle_add(idle_func, user_data);
    }
    idle_count++;
}

void size_notify_init() { adjustment_notify(NULL, NULL); }

void widgetKeyPress(GtkEventControllerKey *self, guint keyval, guint keycode, GdkModifierType state,
                    gpointer user_data);

gboolean key_pressed(GtkEventControllerKey *self, guint keyval, guint keycode, GdkModifierType state,
                     gpointer user_data) {
    widgetKeyPress(self, keyval, keycode, state, user_data);
    return TRUE;
}

void widgetButtonPress(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer user_data);

void button_pressed(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer user_data) {
    widgetButtonPress(self, n_press, x, y, user_data);
}

void widgetButtonRelease(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer user_data);

void button_released(GtkGestureClick *self, gint n_press, gdouble x, gdouble y, gpointer user_data) {
    widgetButtonRelease(self, n_press, x, y, user_data);
}

void widgetMotionNotify(GtkEventControllerMotion *self, gdouble x, gdouble y, gpointer user_data);

void motion_notify(GtkEventControllerMotion *self, gdouble x, gdouble y, gpointer user_data) {
    widgetMotionNotify(self, x, y, user_data);
}

void widgetScroll(GtkEventControllerScroll *self, gdouble dx, gdouble dy, gpointer user_data);

gboolean scroll_notify(GtkEventControllerScroll *self, gdouble dx, gdouble dy, gpointer user_data) {
    widgetScroll(self, dx, dy, user_data);
    return TRUE;
}

/* SimpleAction */

void menuItemActivate(GSimpleAction *action, GVariant *parameter, gpointer data);

void menu_item_activate(GSimpleAction *action, GVariant *parameter, gpointer data) {
    menuItemActivate(action, parameter, data);
}

/* Misc */

void objectWeakRef(gpointer data, GObject *where_the_object_was);

void object_weak_notify(gpointer data, GObject *where_the_object_was) {
    // printf("object_weak_notify: %s %p\n", G_OBJECT_TYPE_NAME(where_the_object_was), where_the_object_was);
    objectWeakRef(data, where_the_object_was);
}
