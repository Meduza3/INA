#[derive(Debug)]
pub struct Tree {
  
  pub root: Option<Box<Node>>,
  height: usize

}

impl Tree {

  pub fn new() -> Self {
    Tree {root: None, height: 0}
  }

  pub fn insert(&mut self, k: i32) {
    //Inserting a new key to the tree
    match &mut self.root {
      None => {
        self.root = Node::new(k).into();
      },
      Some(node) => {
        Tree::insert_recursive(node, k);
      }
    }
  }

    fn insert_recursive(node: &mut Box<Node>, k: i32) {
      if k > node.key {
        match &mut node.right {
          None => {
            node.right = Node::new(k).into();
          },
          Some(node) => {
            Tree::insert_recursive(node, k);
          }
        }
      } else if k < node.key {
        match &mut node.left {
          None => {
            node.left = Node::new(k).into();
          },
          Some(node) => {
            Tree::insert_recursive(node, k);
          }
        }
      } 
    }
  }


  // pub fn delete(k: Node) {
    //TODO: Handle lack of key in the tree

  //}

 // pub fn height() -> i32 {
 //   4
 // }

 // pub fn search(x: Node, k: i32) {
 //   
  //}


#[derive(Debug)]
pub struct Node {
  key: i32,
  parent: Option<Box<Node>>,
  left: Option<Box<Node>>,
  right: Option<Box<Node>>
}

impl Node {
  pub fn new(key: i32) -> Self {
    Node {
      key,
      parent: None,
      left: None,
      right: None
      
    }
  }
}

impl From<Node> for Option<Box<Node>> {
  fn from(node: Node) -> Self {
    Some(Box::new(node))
  }
}
