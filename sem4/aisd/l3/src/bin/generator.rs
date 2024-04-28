use rand::Rng;

fn main() {
    let args: Vec<String> = std::env::args().collect();
    if args.len() != 2{
        eprintln!("🦀: {} <n>", args[0]);
        std::process::exit(1);
    }
    let count: usize = args[1].parse().unwrap();
    print!("{} ", args[1]);
    
    print!("{} ", rand::thread_rng().gen_range(1..args[1].parse().unwrap()));

    for _ in 0..count {
        print!("{} ", rand::thread_rng().gen_range(1..=200));
    }
}