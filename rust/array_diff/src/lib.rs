pub fn array_diff<T: PartialEq + Clone>(a: Vec<T>, b: Vec<T>) -> Vec<T> {
    a.into_iter().filter(|v| !b.contains(v)).collect()
}
