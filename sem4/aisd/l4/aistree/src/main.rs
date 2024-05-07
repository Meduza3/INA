mod bst;

fn main() {
    println!("Hello, world!");
}


#[cfg(test)]
mod tests {
    use crate::bst::Tree;

    use super::*;

    #[test]
    fn tree_building() {
        let mut tree = Tree::new();
        tree.insert(8);
        tree.insert(10);
        tree.insert(3);
        tree.insert(20);

        assert_eq!(tree.root.is_some(), true);
        println!("{:?}", tree);

    }
}