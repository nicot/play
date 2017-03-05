extern crate cocoa;

use self::cocoa::base::{selector, nil, NO, id};
use self::cocoa::foundation::{NSUInteger, NSRect, NSPoint, NSSize, NSAutoreleasePool,
                              NSProcessInfo, NSString};
use self::cocoa::appkit::{NSApp, NSApplication, NSApplicationActivationPolicyRegular, NSWindow,
                          NSTitledWindowMask, NSBackingStoreBuffered, NSMenu, NSMenuItem,
                          NSRunningApplication, NSApplicationActivateIgnoringOtherApps,
                          NSWindowButton};

pub fn make_window() {
    unsafe {
        let _pool = NSAutoreleasePool::new(nil);
        let app = NSApp();
        app.setActivationPolicy_(NSApplicationActivationPolicyRegular);
        create_menu(app);

        // create Window
        let window = NSWindow::alloc(nil)
            .initWithContentRect_styleMask_backing_defer_(NSRect::new(NSPoint::new(0., 0.),
                                                                      NSSize::new(900., 900.)),
                                                          NSTitledWindowMask as NSUInteger,
                                                          NSBackingStoreBuffered,
                                                          NO)
            .autorelease();
        window.cascadeTopLeftFromPoint_(NSPoint::new(20., 20.));
        window.center();

        let button = window.standardWindowButton_(NSWindowButton::NSWindowCloseButton);

        set_title(window);
        window.makeKeyAndOrderFront_(nil);
        let current_app = NSRunningApplication::currentApplication(nil);
        current_app.activateWithOptions_(NSApplicationActivateIgnoringOtherApps);
        app.run();
    }
}

unsafe fn set_title(window: id) {
    let title = NSString::alloc(nil).init_str("Hello World!");
    window.setTitle_(title);
}

// create Application menu
unsafe fn make_app_menu() -> id {
    let app_menu = NSMenu::new(nil).autorelease();
    let quit_prefix = NSString::alloc(nil).init_str("Quit ");
    let quit_title =
        quit_prefix.stringByAppendingString_(NSProcessInfo::processInfo(nil).processName());
    let quit_action = selector("terminate:");
    let quit_key = NSString::alloc(nil).init_str("q");
    let quit_item = NSMenuItem::alloc(nil)
        .initWithTitle_action_keyEquivalent_(quit_title, quit_action, quit_key)
        .autorelease();
    app_menu.addItem_(quit_item);
    app_menu
}

unsafe fn create_menu(app: id) {
    let menubar = NSMenu::new(nil).autorelease();
    let app_menu_item = NSMenuItem::new(nil).autorelease();
    menubar.addItem_(app_menu_item);
    app.setMainMenu_(menubar);
    let app_menu = make_app_menu();
    app_menu_item.setSubmenu_(app_menu);
}
