#[cfg(test)]
mod tests {
    use parity_outlier::find_outlier;

    #[test]
    fn test_find_outlier() {
        let test_data = vec![
            ("test 1", [2,6,8,-10,3], 3),
        ];

        for (desc, input, expect) in test_data {
            let got = find_outlier(&input);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
