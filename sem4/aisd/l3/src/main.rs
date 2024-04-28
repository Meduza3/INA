mod select;
mod randselect;

fn main() {
    println!("Hello, LISTA 3!");
    let mut arr = vec![20, 54, 83, 290, 321, 543, 180, 320, 854, 19, 8, 123, 342, 901];
    let arr_len = arr.len();


    println!("{:?}", &arr);
    println!("{:?}", randselect::rand_select(&mut arr, 0, arr_len - 1, 1));
    println!("{:?}", randselect::rand_select(&mut arr, 0, arr_len - 1, 2));
    println!("{:?}", randselect::rand_select(&mut arr, 0, arr_len - 1, 3));
    println!("{:?}", randselect::rand_select(&mut arr, 0, arr_len - 1, 4));
    println!("{:?}", randselect::rand_select(&mut arr, 0, arr_len - 1, 5));
    

    println!("{:?}", select::select(&mut arr, 0, arr_len - 1, 1));
    println!("{:?}", select::select(&mut arr, 0, arr_len - 1, 2));
    println!("{:?}", select::select(&mut arr, 0, arr_len - 1, 3));
    println!("{:?}", select::select(&mut arr, 0, arr_len - 1, 4));
    println!("{:?}", select::select(&mut arr, 0, arr_len - 1, 5));

}

