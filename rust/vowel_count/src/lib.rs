pub fn get_count(string: &str) -> usize {
    let mut vowels_count: usize = 0;

    for c in string.chars() {
        if ['a', 'e', 'i', 'o', 'u'].iter().any(|t| *t == c) {
            vowels_count+=1
        }
    }

    vowels_count
}
