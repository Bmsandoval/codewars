#[cfg(test)]
mod tests {
    use narcissistic_number::narcissistic;

    #[test]
    fn test_narcissistic() {
        let test_data = vec![
            ("test 1", 7, true),
            ("test 2", 371, true),
            ("test 3", 122, false),
            ("test 4", 4887, false)
        ];

        for (desc, input, expect) in test_data {
            let got = narcissistic(input);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
