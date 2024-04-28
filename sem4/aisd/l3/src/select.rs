pub fn select(arr: &mut [usize], p: usize, q:usize, i:usize) -> usize {
    if p == q {
        return arr[p];
    }

    let pivot_index = partition_median_of_medians(arr, p, q);
    let k = pivot_index - p + 1;

    if i == k {
        arr[pivot_index]
    } else if i < k {
        select(arr, p, pivot_index - 1, i)
    } else {
        select(arr, pivot_index + 1, q, i - k)
    }

}

fn partition_median_of_medians(arr: &mut [usize], p: usize, q: usize) -> usize {
    let mut medians = Vec::new();
    let mut start = p;
    while start <= q {
        let end = std::cmp::min(start + 4, q);
        let median = median_of_five(&mut arr[start..=end]);
        medians.push(median);
        start += 5;
    }

    let medians_len = medians.len();
    let median_of_medians = if medians.len() <= 5 {
        median_of_five(&mut medians) 
    } else {
        select(&mut medians, 0, medians_len - 1, medians_len / 2)
    };

    let mut i = p;
    let mut j = q;
    while i < j {
        while arr[i] < median_of_medians { i += 1; }
        while arr[j] > median_of_medians { j -= 1; }
        if i < j {
            arr.swap(i, j);
        }
    }
    i
}

fn median_of_five(arr: &mut [usize]) -> usize {
    arr.sort();
    arr[arr.len() / 2]
}