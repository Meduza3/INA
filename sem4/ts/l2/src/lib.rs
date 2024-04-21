extern crate petgraph;

pub mod ts {
    
    use petgraph::{dot::{Config, Dot}, graph::{Edge, Node}, stable_graph::StableGraph};


    pub fn print_matrix<T: std::fmt::Display>(matrix: &[Vec<T>]) {
        for row in matrix {
            let formatted_row: Vec<String> = row.iter().map(|item| format!("{: >5}", item)).collect();
            println!("{}", formatted_row.join(" "));
        }
    }


}

