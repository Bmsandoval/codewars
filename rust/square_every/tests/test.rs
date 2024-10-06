#[cfg(test)]
mod tests {
    use square_every::square_digits;

    #[test]
    fn test_square_digits() {
        let test_data = vec![
            ("test 1", 9119, 811181),
            ("test 2", 8191, 641811),
        ];

        for (desc, input, expect) in test_data {
            let got = square_digits(input);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
