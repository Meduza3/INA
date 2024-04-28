use rand::Rng;

pub fn rand_select(arr: &mut Vec<usize>, p: usize, q: usize, i: usize) -> usize {
    if p == q {
        return arr[p as usize];
    }

    let r = rand_partition(arr, p, q); // Globalny indeks pivota
    let k = r - p + 1; //Lokalny indeks pivota

    if i == k {
        return arr[r as usize];
    } else if i <  k{
        return rand_select(arr, p, r - 1, i);
    } else {
        return rand_select(arr, r + 1, q, i - k);
    }
}

pub fn rand_partition(arr: &mut Vec<usize>, p: usize, q: usize) -> usize {
    let pivot_index = p + rand::thread_rng().gen_range(0..=(q - p));
    let pivot = arr[pivot_index];

    arr.swap(pivot_index, q);
    let mut store_index = p;
    for i in p..q {
        if arr[i] < pivot {
            arr.swap(store_index, i);
            store_index += 1;
        }
    }
    arr.swap(store_index, q);
    store_index
}