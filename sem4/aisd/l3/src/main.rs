use std::{fs::OpenOptions, io::{self, BufRead, Write}};

mod select;
mod randselect;
mod select3;
mod select7;
mod select9;

fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args.len() != 2 {
        eprintln!("🦀: {} run <-rand|-select|-select-3>", args[0]); // Use standard error for error messages
        std::process::exit(1); // Exit with a non-zero status code indicating failure
    }

    let mode = match args[1].as_str() {
        "-rand" => "rand",
        "-select" => "select",
        "-select-3" => "select-3",
        "-select-7" => "select-7",
        "-select-9" => "select-9",
        _ => {
            eprintln!("🦀: {} run <-rand|-select|-select-3|-select-7|-select-9>", args[0]); // Uniform error handling
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
        } else if mode == "select" {
            select::select(&mut numbers, 0, size - 1, i, &mut comparisons, &mut swaps)
        } else if mode == "select-3" {
            select3::select(&mut numbers, 0, size - 1, i, &mut comparisons, &mut swaps)
        } else if mode == "select-7" {
            select7::select(&mut numbers, 0, size - 1, i, &mut comparisons, &mut swaps)
        } else if mode == "select-9" {
            select9::select(&mut numbers, 0, size - 1, i, &mut comparisons, &mut swaps)
        } else {
            10
        };
        println!("Porównania: {comparisons}");
        println!("Przestawienia: {swaps}");
        let mut file = OpenOptions::new()
        .append(true)
        .create(true)
        .open(format!("{}.txt", mode)).unwrap();
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

