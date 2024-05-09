#[derive(Debug)]
pub struct Tree {
  
  pub root: Option<Box<Node>>,

}

impl Tree {

  pub fn new() -> Self {
    Tree {root: None}
  }

  pub fn height(&self) -> i32 {
    Tree::height_recursive(&self.root)
  }

  fn height_recursive(node: &Option<Box<Node>>) -> i32 {
    match node {
      None => -1,
      Some(node) => {
        let left_height = Tree::height_recursive(&node.left);
        let right_height = Tree::height_recursive(&node.right);

        1 + left_height.max(right_height)
      }
    }
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
    pub fn delete(&mut self, k: i32) {
      self.root = Tree::delete_recursive(self.root.take(), k);
  }

  fn delete_recursive(node: Option<Box<Node>>, k: i32) -> Option<Box<Node>> {
      match node {
          None => None,
          Some(mut node) => {
              if k < node.key {
                  node.left = Tree::delete_recursive(node.left.take(), k);
              } else if k > node.key {
                  node.right = Tree::delete_recursive(node.right.take(), k);
              } else {
                  // Case 1: Node is a leaf
                  if node.left.is_none() && node.right.is_none() {
                      return None;
                  }
                  // Case 2: Node has only one child
                  else if node.left.is_none() {
                      return node.right.take();
                  } else if node.right.is_none() {
                      return node.left.take();
                  }
                  // Case 3: Node has two children
                  else {
                      let successor = Tree::find_min(node.right.as_mut()).key;
                      node.key = successor;
                      node.right = Tree::delete_recursive(node.right.take(), successor);
                  }
              }
              Some(node)
          }
      }
  }

  fn find_min(node: Option<&mut Box<Node>>) -> &Node {
      if let Some(node) = node {
          if node.left.is_none() {
              node
          } else {
              Tree::find_min(node.left.as_mut())
          }
      } else {
          panic!("Cannot find minimum of empty tree");
      }
  }
    pub fn search(&self, k: i32) -> Option<&Node> {
       match &self.root {
        None => None,
        Some(node) => Tree::search_recursive(node, k)
          
        }
       }
       
    fn search_recursive(node: &Box<Node>, k: i32) -> Option<&Node> {
             if k == node.key {
               Some(node)
             } else if k <= node.key {
               match &node.left {
                 None => None,
                 Some(node) => Tree::search_recursive(node, k)
               }
             } else {
               match &node.right {
                 None => None,
                 Some(node) => Tree::search_recursive(node, k)
               }
           }
    }
}

#[derive(Debug)]
pub struct Node {
  pub key: i32,
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
