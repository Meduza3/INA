use std::io::{self, BufRead};

mod select;
mod randselect;

fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args.len() != 2 {
        eprintln!("🦀: {} run <-rand|-select>", args[0]); // Use standard error for error messages
        std::process::exit(1); // Exit with a non-zero status code indicating failure
    }

    let mode = match args[1].as_str() {
        "-rand" => "rand",
        "-select" => "select",
        _ => {
            eprintln!("🦀: {} run <-rand|-select>", args[0]); // Uniform error handling
            std::process::exit(1);
        }
    };

    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let input = line.unwrap();
        let numbers: Vec<usize> = input
            .split_whitespace()
            .map(|num| num.parse().unwrap())
            .collect();

            if numbers.len() < 3 {
                eprintln!("🦀: Zapomniałeś czegoś w tablicy");
                continue;
            }

        let size = numbers[0];
        let i = numbers[1];
        let mut numbers: Vec<usize> = numbers[2..].to_vec();
        let initial_numbers = numbers.clone();
        let goal = if mode == "rand" {
            randselect::rand_select(&mut numbers, 0, size - 1, i)
        } else {
            select::select(&mut numbers, 0, size - 1, i)
        };
        if size <= 50 {
            println!("Tryb: {mode}");
            println!("Przed SELECT: {:?}", initial_numbers);
            println!("Po SELECT: {:?}", numbers);
            println!("{i}-ta statystyka pozycyjna: {goal}");
            numbers.sort();
            println!("Posortowany: {:?}", numbers);
        }
    }
}

