#[cfg(test)]
mod tests {
    use ordered_unique::unique_in_order;

    #[test]
    fn test_ordered_unique() {
        let char_test_data = vec![
            ("test 1", "AAAABBBCCDAABBB".chars(), vec!['A','B','C','D','A','B']),
            ("test 2", "ABBCcAD".chars(), vec!['A','B','C','c','A','D'] ),
        ];
        let int_test_data = vec![
            ("test 3", vec![1,2,2,3,3], vec![1,2,3] ),
        ];

        for (desc, input, expect) in char_test_data {
            let got = unique_in_order(input);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
        for (desc, input, expect) in int_test_data {
            let got = unique_in_order(input);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
