pub fn unique_in_order<T>(sequence: T) -> Vec<T::Item>
where
    T: std::iter::IntoIterator,
    T::Item: std::cmp::PartialEq + std::fmt::Debug + Clone,
{
    let mut vec: Vec<T::Item> = Vec::new();
    let mut t: Option<T::Item> = None;

    for i in sequence {
        if Some(&i) != t.as_ref() {
            t = Some(i.clone());
            vec.push(i);
        }
    }
    vec
}