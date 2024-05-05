use std::{fs::OpenOptions, io::{self, BufRead, Write}};

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
        let mut comparisons: usize = 0;
        let mut swaps: usize = 0;
        let goal = if mode == "rand" {
            randselect::rand_select(&mut numbers, 0, size - 1, i, &mut comparisons, &mut swaps)
        } else {
            select::select(&mut numbers, 0, size - 1, i, &mut comparisons, &mut swaps)
        };
        println!("Porównania: {comparisons}");
        println!("Przestawienia: {swaps}");
        let mut file = OpenOptions::new()
        .append(true)
        .create(true)
        .open("output.txt").unwrap();
        let data = format!("{}\t {}\t{}\n", size, comparisons, swaps);
        let _ = file.write_all(data.as_bytes());
        println!("Rozmiar: {}\tPorównania: {}\tPrzestawienia: {}", size, comparisons, swaps);
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

