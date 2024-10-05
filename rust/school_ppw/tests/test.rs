#[cfg(test)]
mod tests {
    use school_ppw::paperwork;

    #[test]
    fn test_count_by() {
        let test_data = vec![
            ("test 1", (5,5), 25),
            ("test 2", (5,-5), 0),
            ("test 3", (-5,-5), 0),
            ("test 4", (-5,5), 0),
            ("test 5", (5,0), 0)
        ];

        for (desc, input, expect) in test_data {
            let got = paperwork(input.0, input.1);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
