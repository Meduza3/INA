mod bst;
extern crate rand;
fn main() {
    use crate::bst::Tree;

    println!("Hello, world!");

    let mut tree = Tree::new();
    let mut vector = vec![];
    for _ in 0..50 {  
        let mut rng = rand::thread_rng();
        let random_number = rand::Rng::gen_range(&mut rng, 1..=100);
        vector.push(random_number);
    }
    
    for k in &vector {
        tree.insert(*k);
    }
    println!("{}", tree.height());
    println!("{:?}", tree);

    for k in &vector {
        tree.delete(*k);
    }
    println!("{:?}", tree);
    
    let mut vector_sorted = vector.clone();
    vector_sorted.sort();


    for k in &vector_sorted {
        println!("{}", k);
        tree.insert(*k);
    }
    println!("{}", tree.height());
    println!("{:?}", tree);

    for k in &vector_sorted {
        tree.delete(*k);
    }
    println!("{:?}", tree);
}


#[cfg(test)]
mod tests {
    use crate::bst::Tree;
use rand::Rng;

    use super::*;

    #[test]
    fn tree_building() {
        let mut tree = Tree::new();
        tree.insert(8);
        tree.insert(10);
        tree.insert(3);
        tree.insert(20);
        tree.insert(21);
        tree.insert(1);

        assert_eq!(tree.root.is_some(), true);
        println!("{}", tree.height());
        println!("{:?}", tree);

    }

    #[test]
    fn tree_searching() {
        let mut tree = Tree::new();
        tree.insert(1);
        tree.insert(4);
        tree.insert(3);
        tree.insert(2);
        tree.insert(5);

        tree.search(4).expect("WTF!?");
    }

    #[test]
    fn tree_delete() {
        let mut tree = Tree::new();
        tree.insert(1);
        tree.insert(4);
        tree.insert(3);
        tree.insert(2);
        tree.insert(5);
        println!("{:?}", tree);
        tree.delete(3);
        println!("{:?}", tree);
        assert!(tree.search(3).is_none())
    }
}