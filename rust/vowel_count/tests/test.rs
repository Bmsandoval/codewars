#[cfg(test)]
mod tests {
    use vowel_count::get_count;

    #[test]
    fn test_count_by() {
        let test_data = vec![
            ("test 1", "abracadabra", 5),
            ("test 2", "dolly parton", 3),
            ("test 3", "harry potter", 3),
            ("test 4", "random human a", 5),
            ("test 5", "koda", 2)
        ];

        for (desc, input, expect) in test_data {
            let got = get_count(input);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
