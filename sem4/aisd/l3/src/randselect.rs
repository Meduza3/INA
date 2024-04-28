use rand::Rng;

pub fn rand_select(arr: &mut Vec<usize>, p: usize, q: usize, i: usize, comparisons: &mut usize, swaps: &mut usize) -> usize {
    if p == q {
        return arr[p as usize];
    }

    let r = rand_partition(arr, p, q, comparisons, swaps); // Globalny indeks pivota
    let k = r - p + 1; //Lokalny indeks pivota

    if i == k {
        return arr[r as usize];
    } else if i <  k{
        return rand_select(arr, p, r - 1, i, comparisons, swaps);
    } else {
        return rand_select(arr, r + 1, q, i - k, comparisons, swaps);
    }
}

pub fn rand_partition(arr: &mut Vec<usize>, p: usize, q: usize, comparisons: &mut usize, swaps: &mut usize) -> usize {
    let pivot_index = p + rand::thread_rng().gen_range(0..=(q - p));
    let pivot = arr[pivot_index];

    arr.swap(pivot_index, q);
    *swaps += 1;
    let mut store_index = p;
    *comparisons += 1;
    for i in p..q {
        if arr[i] < pivot {
            arr.swap(store_index, i);
            *swaps += 1;
            store_index += 1;
        }
    }
    arr.swap(store_index, q);
    *swaps += 1;
    store_index
}