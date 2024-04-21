extern crate petgraph;
extern crate rand;
extern crate model_sieci;
extern crate colored;

use std::collections::HashSet;

#[allow(unused_imports)]
use petgraph::stable_graph;
use petgraph::{algo::connected_components, dot::{Config, Dot}, graph::NodeIndex, visit::{Bfs, Dfs, EdgeRef, IntoEdgeReferences, NodeIndexable, Visitable}, Graph, Undirected};
// use petgraph::Undirected;
use rand::Rng;
use colored::*;
use model_sieci::ts;


const BITS_PER_PACKET: u32  = 12000;

// Ilosc bitow w jednym gigabajcie to 1024 * 1024 * 1024 * 8
// 20 gigabajtow dziennie = 8589934592 bitow dziennie
// 357913941.333 bitow na godzine
// 5965232.35556 bitow na minute
// 99420.5392593 bitow na sekunde
// 8 pakietow na sekunde

fn main() {
    let mut n: Vec<Vec<f64>> = vec![vec![0.0; 20]; 20];
    setup_n(&mut n);


    let mut graph = stable_graph::StableGraph::<Node, Edge>::new();
    setup_graph(&mut graph, &n);
    ts::print_matrix(&n);

    //println!("{:?}", graph);

    //Eksperyment niezawodnosci sieci:

    let mut rng =  rand::thread_rng();
    let edge_indices: Vec<_> = graph.edge_indices().collect();

    export_to_dot(&graph);

    for e in edge_indices {
        let oops = rng.gen_range(0_f64..1_f64);
        let niezawodnosc = graph.edge_weight(e).unwrap().niezawodnosc;
        let (source, target) = graph.edge_endpoints(e).unwrap();

        if oops > niezawodnosc {
            graph.remove_edge(e);
            print!("{}", format!("Edge from {} to {} failed!\n", source.index(), target.index()).red());
        } else {
            print!("{}", "Wszystko dobrze!\n".green());
        }
    }

    if is_connected(&graph) {
        println!("{}", "The graph is still connected.".bold().green());
    } else {
        println!("{}", "The graph is not connected.".bold().red());
    }

}

fn export_to_dot(graph: &stable_graph::StableGraph<Node, Edge>) {
    let dot = Dot::with_config(graph, &[Config::EdgeNoLabel]);
    let output = format!("{:?}", dot);
    std::fs::write("graph.dot", output).expect("Unable to write DOT file");
}

fn is_connected(graph: &stable_graph::StableGraph<Node, Edge>) -> bool {
    let mut bfs = Bfs::new(graph, NodeIndex::new(0)); // Start BFS from node index 0
    let mut visited = HashSet::new();

    while let Some(nx) = bfs.next(graph) {
        visited.insert(nx);
        //println!("{:?}", nx)
    }

    visited.len() == graph.node_count()
}


fn setup_n(n: &mut Vec<Vec<f64>>) {
    let mut rng = rand::thread_rng();

    n[0][1] = low_usage(&mut rng);
    n[0][4] = low_usage(&mut rng);
    n[1][2] = low_usage(&mut rng);
    n[2][3] = medium_usage(&mut rng);
    n[2][8] = high_usage(&mut rng);
    n[3][4] = low_usage(&mut rng);
    n[3][17] = high_usage(&mut rng);
    n[5][6] = low_usage(&mut rng);
    n[5][9] = low_usage(&mut rng);
    n[6][7] = low_usage(&mut rng);
    n[7][8] = medium_usage(&mut rng);
    n[7][13] = high_usage(&mut rng);
    n[8][9] = low_usage(&mut rng);
    n[10][11] = low_usage(&mut rng);
    n[10][14] = low_usage(&mut rng);
    n[11][12] = low_usage(&mut rng);
    n[12][13] = medium_usage(&mut rng);
    n[12][18] = high_usage(&mut rng);
    n[13][14] = low_usage(&mut rng);
    n[15][16] = low_usage(&mut rng);
    n[15][19] = low_usage(&mut rng);
    n[16][17] = low_usage(&mut rng);
    n[17][18] = medium_usage(&mut rng);
    n[18][19] = low_usage(&mut rng);

    n[1][0] = n[0][1];
    n[4][0] = n[0][4];
    n[2][1] = n[1][2];
    n[3][2] = n[2][3];
    n[8][2] = n[2][8];
    n[4][3] = n[3][4];
    n[17][3] = n[3][17];
    n[6][5] = n[5][6];
    n[9][5] = n[5][9];
    n[7][6] = n[6][7];
    n[8][7] = n[7][8];
    n[13][7] = n[7][13];
    n[9][8] = n[8][9];
    n[11][10] = n[10][11];
    n[14][10] = n[10][14];
    n[12][11] = n[11][12];
    n[13][12] = n[12][13];
    n[18][12] = n[12][18];
    n[14][13] = n[13][14];
    n[16][15] = n[15][16];
    n[19][15] = n[15][19];
    n[17][16] = n[16][17];
    n[18][17] = n[17][18];
    n[19][18] = n[18][19];

}

fn low_usage(rng: &mut impl Rng) -> f64 {
    return (rng.gen_range(7.0_f64..9.0) * 100.0).round() / 100.0;
}

fn medium_usage(rng: &mut impl Rng) -> f64 {
    return (rng.gen_range(14.0_f64..18.0_f64) * 100.0).round() / 100.0;
}

fn high_usage(rng: &mut impl Rng) -> f64 {
    return (rng.gen_range(28.0_f64..36.0_f64) * 100.0).round() / 100.0;
}


fn setup_graph(graph: &mut stable_graph::StableGraph<Node, Edge>, n: &Vec<Vec<f64>>) {
    for _ in 1..21 {
        graph.add_node(Node::new());
        }
        graph.add_edge(stable_graph::node_index(0), stable_graph::node_index(1), Edge::new(1, 0.95, 10 * BITS_PER_PACKET, *n.get(0).unwrap().get(1).unwrap()));
        graph.add_edge(stable_graph::node_index(1), stable_graph::node_index(2), Edge::new(1, 0.95,10 * BITS_PER_PACKET, *n.get(1).unwrap().get(2).unwrap()));
        graph.add_edge(stable_graph::node_index(2), stable_graph::node_index(3), Edge::new(1, 0.97,20 * BITS_PER_PACKET, *n.get(2).unwrap().get(3).unwrap()));
        graph.add_edge(stable_graph::node_index(3), stable_graph::node_index(4), Edge::new(1, 0.95, 10 * BITS_PER_PACKET, *n.get(3).unwrap().get(4).unwrap()));
        graph.add_edge(stable_graph::node_index(4), stable_graph::node_index(0), Edge::new(1, 0.95, 10 * BITS_PER_PACKET, *n.get(4).unwrap().get(0).unwrap()));
        
        graph.add_edge(stable_graph::node_index(5), stable_graph::node_index(6), Edge::new(1, 0.95,10 * BITS_PER_PACKET, *n.get(5).unwrap().get(6).unwrap()));
        graph.add_edge(stable_graph::node_index(6), stable_graph::node_index(7), Edge::new(1, 0.95,10 * BITS_PER_PACKET, *n.get(6).unwrap().get(7).unwrap()));
        graph.add_edge(stable_graph::node_index(7), stable_graph::node_index(8), Edge::new(1, 0.97,20 * BITS_PER_PACKET, *n.get(7).unwrap().get(8).unwrap()));
        graph.add_edge(stable_graph::node_index(8), stable_graph::node_index(9), Edge::new(1, 0.95,10 * BITS_PER_PACKET, *n.get(8).unwrap().get(9).unwrap()));
        graph.add_edge(stable_graph::node_index(9), stable_graph::node_index(5), Edge::new(1, 0.95, 10 * BITS_PER_PACKET, *n.get(9).unwrap().get(5).unwrap()));

        graph.add_edge(stable_graph::node_index(10), stable_graph::node_index(11), Edge::new(1,0.95, 10 * BITS_PER_PACKET, *n.get(10).unwrap().get(11).unwrap()));
        graph.add_edge(stable_graph::node_index(11), stable_graph::node_index(12), Edge::new(1,0.95, 10 * BITS_PER_PACKET, *n.get(11).unwrap().get(12).unwrap()));
        graph.add_edge(stable_graph::node_index(12), stable_graph::node_index(13), Edge::new(1,0.97, 20 * BITS_PER_PACKET, *n.get(12).unwrap().get(13).unwrap()));
        graph.add_edge(stable_graph::node_index(13), stable_graph::node_index(14), Edge::new(1,0.95, 10 * BITS_PER_PACKET, *n.get(13).unwrap().get(14).unwrap()));
        graph.add_edge(stable_graph::node_index(14), stable_graph::node_index(10), Edge::new(1, 0.95, 10 * BITS_PER_PACKET, *n.get(14).unwrap().get(10).unwrap()));
    
        graph.add_edge(stable_graph::node_index(15), stable_graph::node_index(16), Edge::new(1,0.95, 10 * BITS_PER_PACKET, *n.get(15).unwrap().get(16).unwrap()));
        graph.add_edge(stable_graph::node_index(16), stable_graph::node_index(17), Edge::new(1, 0.95,10 * BITS_PER_PACKET, *n.get(16).unwrap().get(17).unwrap()));
        graph.add_edge(stable_graph::node_index(17), stable_graph::node_index(18), Edge::new(1,0.97, 20 * BITS_PER_PACKET, *n.get(17).unwrap().get(18).unwrap()));
        graph.add_edge(stable_graph::node_index(18), stable_graph::node_index(19), Edge::new(1, 0.95,10 * BITS_PER_PACKET, *n.get(18).unwrap().get(19).unwrap()));
        graph.add_edge(stable_graph::node_index(19), stable_graph::node_index(15), Edge::new(1, 0.95, 10 * BITS_PER_PACKET, *n.get(19).unwrap().get(15).unwrap()));
    
        graph.add_edge(stable_graph::node_index(2), stable_graph::node_index(8), Edge::new(1, 0.99,40 * BITS_PER_PACKET, *n.get(2).unwrap().get(8).unwrap()));
        graph.add_edge(stable_graph::node_index(7), stable_graph::node_index(13), Edge::new(1, 0.99,40 * BITS_PER_PACKET, *n.get(7).unwrap().get(13).unwrap()));
        graph.add_edge(stable_graph::node_index(12), stable_graph::node_index(18), Edge::new(1,0.99, 40 * BITS_PER_PACKET, *n.get(12).unwrap().get(18).unwrap()));
        graph.add_edge(stable_graph::node_index(3), stable_graph::node_index(17), Edge::new(1, 0.99,40 * BITS_PER_PACKET, *n.get(3).unwrap().get(17).unwrap()));
}

#[derive(Default, Debug)]
struct Node {
}

impl Node {
    fn new() -> Node {
        Node {}
    }
}
 
#[derive(Debug)]
 struct Edge {
    index: u8,
    niezawodnosc: f64,
    c_capacity: u32, //Maksymalna liczba bitow, ktora mozna wprowadzic do kanalu komunikacyjnego w ciegu sekundy
    a_flow: f64,  //Faktyczna liczba pakietow, ktore wprowadza sie do kanalu komunikacyjnego w ciagu sekundy
    enabled: bool
 }



impl Edge {
    fn capacity(&self) -> &u32 {
        &self.c_capacity
    }

    fn a_flow(&self) -> &f64 {
        &self.a_flow
    }

    fn a_flow_in_bits(&self) -> f64 {
        self.a_flow * f64::from(BITS_PER_PACKET) 
    }

    fn new(index: u8, niezawodnosc: f64, capacity: u32, a_flow: f64) -> Edge {
        Edge {
            index: index,
            niezawodnosc: niezawodnosc,
            c_capacity: capacity,
            a_flow: a_flow,
            enabled: true
        }
    }

}


// S = <G, H> - Model sieci, G to graf a H to zbiór funkcji przyporządkujących łukom wagi
// N=[n(i,j)] - Macierz natężeń strumienia pakietow
// element n(i,j) - Liczba pakietów przesyłanych w ciągu sekunda od v(i) do v(j)
 
//Topologia grafu G tak, zeby zaden wierzholek nie byl izolowany oraz
 
// Funkcja przepustowoci c (maksymalna liczba bitow, ktora mozna wprowadzic do kanalu komunikacyjnego w ciagu sekundy)
// Funkcja przeplywu a (faktyczna liczba pakietow, ktore wprowadza sie do kanalu komunikacyjnego w ciagu sekundy)
// for all e z E: c(e) > a(e)