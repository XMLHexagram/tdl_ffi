extern crate libc;
use std::ffi::{CStr, CString};
use std::{thread, time::Duration};

use std::net::Shutdown::Read;

#[link(name = "tdl", kind = "static" )]
extern {
    fn Test() -> CString;
    fn StartGrpcServer(port: libc::c_int) -> CString;
}

fn main() {
    unsafe {
        let a = Test();
        println!("{}", a.to_str().unwrap());
        let b = StartGrpcServer(8091);
        println!("{}", b.to_str().unwrap());
        // thread::sleep(Duration::from_millis(4000));
        // let b = CString::new("tdl ttt").expect("114514");
        // println!("{}", a.to_str().unwrap());

        // let b = CString::new("tdl ttt").expect("114514");
        // println!("{}", a.to_str().unwrap());
        // // RunScript(b.clone());
        // // RunScript(b.clone());
        // let c = ReadLine();
        // println!("{}", c.to_str().unwrap());
        // Cancel();
    }
}
