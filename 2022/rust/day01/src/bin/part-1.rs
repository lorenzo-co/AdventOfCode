use day01::process_part1;
use std::fs;
use std::env;

fn main() {
    env::set_var("RUST_BACKTRACE", "1");
    let file = fs::read_to_string("./input.txt").unwrap();
    println!("{}", process_part1(&file));
}