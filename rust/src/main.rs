use std::rand;

fn main() {
}

#[test]
fn qsort_test() {
    let mut l = vec![3i, 2, 1, 0];
    qsort(l.as_mut_slice());
    assert!(l == [0, 1, 2, 3]);
}

fn qsort<T: Ord>(l: &mut [T]) {
    let length = l.len();
    if length > 1 {
        let pivot_index = partition(l);
        qsort(l.slice_mut(0, pivot_index-1));
        qsort(l.slice_mut(pivot_index, length));
    }

    fn partition<T: Ord>(l: &mut [T]) -> uint {
        let length = l.len();
        let pivot_index = rand::random::<uint>()%(length-1);
        l.swap(pivot_index, length-1);

        let mut store_index = 0;
        let mut i = 0u;
        while i < length {
            if l[i] <= l[length-1] {
                l.swap(i, store_index);
                store_index += 1;
            }
            i += 1;
        }
        store_index
    }
}

