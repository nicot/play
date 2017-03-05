extern crate gl;
extern crate glutin;
extern crate libc;

pub fn make_window() {
    let window = glutin::Window::new().unwrap();

    unsafe { window.make_current() };

    unsafe {
        gl::load_with(|symbol| window.get_proc_address(symbol) as *const _);

        gl::ClearColor(0.0, 0.0, 0.0, 1.0);
    }

    for event in window.wait_events() {
        unsafe {
			gl::Clear(gl::COLOR_BUFFER_BIT)
		};
        window.swap_buffers();

        match event {
            glutin::Event::Closed => break,
            _ => ()
        }
    }
}
