extern crate petgraph;
extern crate colored;

use std::{fs::File, io::Write};

use petgraph::{algo::connected_components, graph::{Node, NodeIndex, UnGraph}, prelude, visit::IntoEdges};
use colored::*;
use rand::{rngs::ThreadRng, Rng};
use petgraph::algo::astar;
const BITS_PER_PACKET: i32 = 12000;

fn main() {

    let mut graph: UnGraph<i32, Edge> = UnGraph::new_undirected();

    let mut nodes = Vec::new();
    for _ in 0..20 {
        nodes.push(graph.add_node(0));
    }

    let mut rng = rand::thread_rng();
    let mut n: Vec<Vec<f64>> = vec![];
    for _ in 0..20 {
        let mut row: Vec<f64> = vec![];
        for _ in 0..20 {
            row.push(0.0);
        }
        n.push(row);
    }
    for i in 0..20 {
        for j in 0..20 {
            if i != j {
                n[i][j] = rng.gen_range(1.0..5.0);
            } else {
                n[i][j] = 0.0;
            }
        }
    }

    //println!("{:?}", n);

    graph.add_edge(nodes[0], nodes[1], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[1], nodes[2], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[2], nodes[3], Edge::new(700 * BITS_PER_PACKET, 0.95));
    graph.add_edge(nodes[3], nodes[4], Edge::new(400 * BITS_PER_PACKET,  0.90));
    graph.add_edge(nodes[4], nodes[0], Edge::new(400 * BITS_PER_PACKET, 0.90));

    graph.add_edge(nodes[5], nodes[6], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[6], nodes[7], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[7], nodes[8], Edge::new(700 * BITS_PER_PACKET, 0.95));
    graph.add_edge(nodes[8], nodes[9], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[9], nodes[5], Edge::new(400 * BITS_PER_PACKET, 0.90));

    graph.add_edge(nodes[10], nodes[11], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[11], nodes[12], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[12], nodes[13], Edge::new(700 * BITS_PER_PACKET, 0.95));
    graph.add_edge(nodes[13], nodes[14], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[14], nodes[10], Edge::new(400 * BITS_PER_PACKET, 0.90));

    graph.add_edge(nodes[15], nodes[16], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[16], nodes[17], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[17], nodes[18], Edge::new(700 * BITS_PER_PACKET, 0.95));
    graph.add_edge(nodes[18], nodes[19], Edge::new(400 * BITS_PER_PACKET, 0.90));
    graph.add_edge(nodes[19], nodes[15], Edge::new(400 * BITS_PER_PACKET, 0.90));

    graph.add_edge(nodes[2], nodes[8], Edge::new(1000 * BITS_PER_PACKET, 0.98));
    graph.add_edge(nodes[7], nodes[13], Edge::new(1000 * BITS_PER_PACKET, 0.98));
    graph.add_edge(nodes[12], nodes[18], Edge::new(1000 * BITS_PER_PACKET, 0.98));
    graph.add_edge(nodes[3],nodes[17], Edge::new(1000 * BITS_PER_PACKET,  0.98));

    println!("{}", "Eksperyment 1:".underline());
    run_experiment1(&mut graph.clone(), &mut n.clone(), &mut rng.clone(), &nodes.clone());

    println!("{}", "Eksperyment 2:".underline());
    run_experiment2(&mut graph.clone(), &mut n.clone(), &mut rng.clone(), &nodes.clone());

    println!("{}", "Eksperyment 3:".underline());
    run_experiment3(&mut graph.clone(), &mut n.clone(), &mut rng.clone(), &nodes.clone());
    
}

fn run_experiment1(graph: &mut UnGraph<i32, Edge>, n: &mut Vec<Vec<f64>>, rng: &mut ThreadRng, nodes: &Vec<NodeIndex>) {
    let mut file = File::create("data1.txt").unwrap();

    for _ in 0..10 {
        let (s, n_c, o, t_s) = run_simulation(graph.clone(), n.clone(), rng.clone(), nodes.clone(), 0.3, 100);

        println!("{}", format!("Success: {}% | Not connected: {}%, Overloaded: {}%, Too slow: {}%", s as f64, n_c as f64, o as f64, t_s as f64).bright_purple().bold());
        file.write_all(format!("{}  {}  {}  {}\n", s, n_c, o, t_s).as_bytes());
        println!("Increasing n[*][*] by *= 1.1");
        increase_values(n, 1.1);
    }
}

fn run_experiment2(graph: &mut UnGraph<i32, Edge>, n: &mut Vec<Vec<f64>>, rng: &mut ThreadRng, nodes: &Vec<NodeIndex>) {
    let mut file = File::create("data2.txt").unwrap();

    for _ in 0..10 {
        let (s, n_c, o, t_s) = run_simulation(graph.clone(), n.clone(), rng.clone(), nodes.clone(), 0.3, 100);

        println!("{}", format!("Success: {}% | Not connected: {}%, Overloaded: {}%, Too slow: {}%", s as f64, n_c as f64, o as f64, t_s as f64).bright_purple().bold());
        file.write_all(format!("{}  {}  {}  {}\n", s, n_c, o, t_s).as_bytes());
        println!("Increasing capacities by *= 1.1");
        increase_capacity(graph, 1.1);
    }
}

fn run_experiment3(mut graph: &mut UnGraph<i32, Edge>, n: &mut Vec<Vec<f64>>, rng: &mut ThreadRng, nodes: &Vec<NodeIndex>) {
    let mut file = File::create("data3.txt").unwrap();

    for _ in 0..10 {
        let (s, n_c, o, t_s) = run_simulation(graph.clone(), n.clone(), rng.clone(), nodes.clone(), 0.3, 100);

        println!("{}", format!("Success: {}% | Not connected: {}%, Overloaded: {}%, Too slow: {}%", s as f64, n_c as f64, o as f64, t_s as f64).bright_purple().bold());
        file.write_all(format!("{}  {}  {}  {}\n", s, n_c, o, t_s).as_bytes());

        add_random_edge(&mut graph);
    }
}

fn increase_capacity(graph: &mut UnGraph<i32, Edge>, increase: f64) {
    for edge_index in graph.edge_indices() {
        if let Some(edge) = graph.edge_weight_mut(edge_index) {
            edge.capacity = (edge.capacity as f64 * increase) as i32;
        }
    }
}

fn increase_values(n: &mut Vec<Vec<f64>>, increase: f64) {
    for i in 0..n.len() {
        for j in 0..n[i].len() {
            n[i][j] *= increase;
        }
    }
}

fn add_random_edge(graph: &mut UnGraph<i32, Edge>) {
    //Add a random edge between two existing nodes
    let mut rng = rand::thread_rng();
    let num_nodes = graph.node_count();
    if num_nodes >= 2 {
        let node1_index = rng.gen_range(0..num_nodes);
        let node2_index = loop {
            let index = rng.gen_range(0..num_nodes);
            if index != node1_index {
                break index;
            }
        };
        let node1 = NodeIndex::new(node1_index);
        let node2 = NodeIndex::new(node2_index);
        let capacity = rng.gen_range(200..1000) * BITS_PER_PACKET;
        let reliability = rng.gen_range(0.8..1.0);
        println!("Added between {:?} and {:?}", node1, node2);
        graph.add_edge(node1, node2, Edge::new(capacity, reliability));
    }
}

fn run_simulation(graph: UnGraph<i32, Edge>, n: Vec<Vec<f64>>, rng: ThreadRng, nodes: Vec<NodeIndex>, threshold: f64, iterations: i32) -> (i32, i32, i32, i32) {
    let mut success = 0;
    let mut not_connected = 0;
    let mut overloaded = 0;
    let mut too_slow = 0;

    for _ in 0..iterations {
        match perform_experiment(graph.clone(), n.clone(), rng.clone(), nodes.clone(), threshold) {
            Ok(_) => {
                success = success + 1
            },
            Err(e) => {
                match e {
                    FailedExperiment::NotConnected => not_connected += 1,
                    FailedExperiment::Overloaded => overloaded += 1,
                    FailedExperiment::TooSlow => too_slow += 1
                }
            }
        }
    }

    return (success, not_connected, overloaded, too_slow);
}

enum FailedExperiment {
    NotConnected,
    Overloaded,
    TooSlow,
}

fn perform_experiment(mut graph: UnGraph<i32, Edge>, n: Vec<Vec<f64>>, mut rng: ThreadRng, nodes: Vec<NodeIndex>, threshold: f64) -> Result<bool, FailedExperiment> {

    for edge_index in graph.edge_indices() {
        if let Some(edge) = graph.edge_weight(edge_index) {
            let oops = rng.gen_range(0.0..1.0);
            if let Some((source, target)) = graph.edge_endpoints(edge_index) {
                if oops > edge.reliability {
                    //println!("{}", format!("Łącze między {} a {}: Wyjebało", source.index(), target.index()).red().bold());
                    graph.remove_edge(edge_index);
                } else {
                    //println!("{}", format!("Łącze między {} a {}: OK", source.index(), target.index()).green().bold());
                }
            }
        }
    }
    assign_flow(&mut graph, n.clone(), nodes);

    //Test 1: Czy wszystkie ośrodki połączone?.
    if connected_components(&graph) == 1 {
        //println!("{}", "Wszystkie ośrodki na wyspach połączone!".yellow().bold());
    } else {
        //println!("{}", "Nie wszystkie ośrodki połączone! Eksperyment nieudany..".purple().bold());
        return Err(FailedExperiment::NotConnected);
    }

    // Test 2: Sprawdz, czy zawsze flow < capacity
    let mut loss_sum = 0;
    for edge_index in graph.edge_indices() {
        if let Some(edge) = graph.edge_weight_mut(edge_index) {
            if edge.flow * BITS_PER_PACKET > edge.capacity {
                let loss = (edge.flow * BITS_PER_PACKET - edge.capacity)/ BITS_PER_PACKET;
                loss_sum += loss;
                //println!("{}", format!("Łącze {:?} przeciążone o {} pakietów", edge_index, loss).purple().bold());
            } else {
                //println!("{}", format!("Flow NOT exceeds capacity on edge {:?}", edge_index).yellow().bold());
            }
        }
    }
    if loss_sum != 0 {
        //println!("{}", format!("Łacznie {} pakietów ponad miarę. Eksperyment nieudany", loss_sum).purple().bold(),);
        return Err(FailedExperiment::Overloaded);
    }


    //Test 3: Jak tam predkosc?
    let sum: f64 = n.iter().flatten().map(|&x| x as f64).sum();
    //println!("Sum of all cells: {}", sum);
    let mut middle = 0.0;
    for edge_index in graph.edge_indices() {
        if let Some(edge) = graph.edge_weight(edge_index) {
            middle += edge.flow as f64 / (edge.capacity as f64 / BITS_PER_PACKET as f64 - edge.flow as f64);
        }
    }
    let t_sr = 1.0 / sum * middle;
    //println!("{}", format!("T_sr = {}", t_sr).bold());
    if t_sr > threshold {
        //println!("{}", "Mega dlugo zajmuje to. Eksperyment nieudany".purple().bold());
        return Err(FailedExperiment::TooSlow);
    }

    //println!("{}", "Wszystko się udało!".yellow().bold());
    return Ok(true);

    //println!("{:?}", graph)
}

fn assign_flow(graph: &mut UnGraph<i32, Edge>, n: Vec<Vec<f64>>,  nodes: Vec<NodeIndex>) {
    for i in 0..20 {
        for j in 0..20 {
            if i != j {
                let n_value = n[i][j] as i32; // Convert n_value to i32
                let path = astar(&*graph, nodes[i], |finish| finish == nodes[j], |_| 0, |_| 0);
                if let Some((_, path)) = path {
                    for window in path.windows(2) {
                        if let [node1, node2] = *window {
                            if let Some(edge) = graph.find_edge(node1, node2) {
                                let edge = graph.edge_weight_mut(edge).unwrap();
                                edge.flow += n_value;
                            }
                        }
                    }
                }
            }
        }
    }
}


#[derive(Debug, Clone)] // Implement the Clone trait for the Edge struct
struct Edge {
    capacity: i32,
    flow: i32,
    reliability: f64,
}

impl Edge {
    fn new(capacity: i32, reliability: f64) -> Edge {
        Edge {capacity, flow: 0, reliability}
    }
}