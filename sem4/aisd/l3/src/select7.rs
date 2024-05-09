pub fn select(arr: &mut [usize], p: usize, q:usize, i:usize, comparisons: &mut usize, swaps: &mut usize) -> usize {
    if p == q {
        return arr[p];
    }

    let pivot_index = partition_median_of_medians(arr, p, q, comparisons, swaps);
    let k = pivot_index - p + 1;

    if arr.len() <= 50 {
        println!("{:?}", arr);
    }

    if i == k {
        arr[pivot_index]
    } else if i < k {
        select(arr, p, pivot_index - 1, i, comparisons, swaps)
    } else {
        select(arr, pivot_index + 1, q, i - k, comparisons, swaps)
    }

}

fn partition_median_of_medians(arr: &mut [usize], p: usize, q: usize, comparisons: &mut usize, swaps: &mut usize) -> usize {
    let mut medians = Vec::new();
    let mut start = p;
    while start <= q {
        let end = std::cmp::min(start + 6, q);
        let median = median_of_seven(&mut arr[start..=end], comparisons, swaps);
        medians.push(median);
        start += 5;
    }

    let medians_len = medians.len();
    let median_of_medians = if medians.len() <= 7 {
        median_of_seven(&mut medians, comparisons, swaps) 
    } else {
        select(&mut medians, 0, medians_len - 1, medians_len / 2, comparisons, swaps)
    };

    let pivot_index = arr.iter().position(|&x| x == median_of_medians).unwrap();
    
    arr.swap(pivot_index, q);
    let mut store_index = p;

    for i in p..q {
        *comparisons += 1;
        if arr[i] < median_of_medians {
            arr.swap(store_index, i);
            *swaps += 1;
            store_index += 1;
        }
    }
    
    arr.swap(store_index, q);
    store_index
}

fn median_of_seven(arr: &mut [usize], comparisons: &mut usize, swaps: &mut usize) -> usize {

    for i in 0..arr.len() {
        for j in i + 1..arr.len() {
            *comparisons += 1;
            if arr[i] > arr[j] {
                arr.swap(i, j);
                *swaps += 1;
            }
        }
    }

    arr[arr.len() / 2]
}