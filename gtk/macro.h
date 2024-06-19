#ifndef _MACRO_H_
#define _MACRO_H_

#include <gtk/gtk.h>

inline GApplication *appToGApplication(GtkApplication *app) { return G_APPLICATION(app); }

inline GObject *appToGObject(GtkApplication *app) { return G_OBJECT(app); }
inline GObject *widgetToGObject(GtkWidget *widget) { return G_OBJECT(widget); }
inline GObject *simpleToGObject(GSimpleAction *simpleAction) { return G_OBJECT(simpleAction); }

inline gpointer appToGPointer(GtkApplication *app) { return app; }
inline gpointer widgetToGPointer(GtkWidget *widget) { return widget; }
inline gpointer layoutToGPointer(PangoLayout *layout) { return layout; }
inline gpointer menuToGPointer(GMenu *menu) { return menu; }
inline gpointer menuItemToGPointer(GMenuItem *item) { return item; }

inline GAction *simpleToGAction(GSimpleAction *simpleAction) { return G_ACTION(simpleAction); }

inline GActionMap *appToGActionMap(GtkApplication *app) { return G_ACTION_MAP(app); }

inline GMenuModel *menuToGMenuModel(GMenu *menu) { return G_MENU_MODEL(menu); }

inline GtkWindow *widgetToGtkWindow(GtkWidget *widget) { return GTK_WINDOW(widget); }

inline GtkContainer *widgetToGtkContainer(GtkWidget *widget) { return GTK_CONTAINER(widget); }

inline GtkLayout *widgetToGtkLayout(GtkWidget *widget) { return GTK_LAYOUT(widget); }
inline gboolean widgetIsLayout(GtkWidget *widget) { return GTK_IS_LAYOUT(widget); }

inline GtkFixed *widgetToGtkFixed(GtkWidget *widget) { return GTK_FIXED(widget); }

inline gulong GSignalConnect(GObject *instance, const gchar *detailed_signal, GCallback c_handler, gpointer data) {
    return g_signal_connect_data(instance, detailed_signal, c_handler, data, NULL, 0);
}

#endif
