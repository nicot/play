#![feature(phase)]
#[phase(plugin, link)] extern crate log;

use std::io::{TcpListener, TcpStream};
use std::io::{Acceptor, Listener};

fn handle_client(mut stream: TcpStream) {
    match stream.write(b"HELLO!") {
        Ok(()) => info!("Response sent"),
        Err(e) => warn!("Failed to respond {}", e),
    }
    let mut buf = [100];
    let n = stream.read(&mut buf);
    println!("{}: {}", n, buf);
    drop(stream);
}

fn main() {

    let listener = TcpListener::bind("127.0.0.1:1080").unwrap();

    // bind the listener to the specified address
    let mut acceptor = listener.listen();

    // accept connections and process them, spawning a new tasks for each one
    for stream in acceptor.incoming() {
        match stream {
            Err(e) => { warn!("fail: {}", e); }
            Ok(stream) => spawn(proc() {
                // connection succeeded
                info!("Successfully accepted TCP connection.");
                handle_client(stream)
            })
        }
    }

    // close the socket server
    drop(acceptor);
}
