#ifndef _MACRO_H_
#define _MACRO_H_

#include <gtk/gtk.h>

inline GApplication *applicationToGApplication(GtkApplication *app) { return G_APPLICATION(app); }

inline GObject *applicationToGObject(GtkApplication *app) { return G_OBJECT(app); }
inline GObject *widgetToGObject(GtkWidget *widget) { return G_OBJECT(widget); }
inline GObject *adjustmentToGObject(GtkAdjustment *adjustment) { return G_OBJECT(adjustment); }
inline GObject *simpleToGObject(GSimpleAction *simpleAction) { return G_OBJECT(simpleAction); }
inline GObject *controllerToGObject(GtkEventController *c) { return G_OBJECT(c); }
inline GObject *gestureToGObject(GtkGesture *c) { return G_OBJECT(c); }

inline gpointer applicationToGPointer(GtkApplication *app) { return app; }
inline gpointer widgetToGPointer(GtkWidget *widget) { return widget; }
inline gpointer layoutToGPointer(PangoLayout *layout) { return layout; }
inline gpointer menuToGPointer(GMenu *menu) { return menu; }
inline gpointer menuItemToGPointer(GMenuItem *item) { return item; }

inline GAction *simpleToGAction(GSimpleAction *simpleAction) { return G_ACTION(simpleAction); }

inline GActionMap *applicationToGActionMap(GtkApplication *app) { return G_ACTION_MAP(app); }

inline GMenuModel *menuToGMenuModel(GMenu *menu) { return G_MENU_MODEL(menu); }

inline GtkWindow *widgetToGtkWindow(GtkWidget *widget) { return GTK_WINDOW(widget); }

inline GtkApplicationWindow *widgetToGtkApplicationWindow(GtkWidget *widget) { return GTK_APPLICATION_WINDOW(widget); }

inline GtkScrolledWindow *widgetToGtkScrolledWindow(GtkWidget *widget) { return GTK_SCROLLED_WINDOW(widget); }

inline GtkFixed *widgetToGtkFixed(GtkWidget *widget) { return GTK_FIXED(widget); }

inline GtkDrawingArea *widgetToGtkDrawingArea(GtkWidget *widget) { return GTK_DRAWING_AREA(widget); }

inline GtkEventController *keyToEventController(GtkEventControllerKey *c) { return GTK_EVENT_CONTROLLER(c); }
inline GtkEventController *gestureToEventController(GtkGesture *c) { return GTK_EVENT_CONTROLLER(c); }
inline GtkEventController *clickToEventController(GtkGestureClick *c) { return GTK_EVENT_CONTROLLER(c); }
inline GtkEventController *scrollToEventController(GtkEventControllerScroll *c) { return GTK_EVENT_CONTROLLER(c); }
inline GtkEventController *motionToEventController(GtkEventControllerMotion *c) { return GTK_EVENT_CONTROLLER(c); }

inline GtkGestureSingle *gestureToGestureSingle(GtkGesture *c) { return GTK_GESTURE_SINGLE(c); }

inline GdkClipboard *objectToGdkClipboard(GObject *o) { return GDK_CLIPBOARD(o); }

inline gulong GSignalConnect(GObject *instance, const gchar *detailed_signal, GCallback c_handler, gpointer data) {
    return g_signal_connect_data(instance, detailed_signal, c_handler, data, NULL, 0);
}

inline void GSignalHandlersDisconnectByFunc(GObject *instance, GCallback c_handler, gpointer data) {
    g_signal_handlers_disconnect_by_func(instance, c_handler, data);
}

#endif
