#[cfg(test)]
mod tests {
    use multiples_of_x::count_by;

    #[test]
    fn test_banjo() {
        let test_data = vec![
            ("test 1", vec![1,2,3,4,5,6,7,8,9,10], (1, 10)),
            ("test 2", vec![2,4,6,8,10], (2, 5)),
            ("test 3", vec![3,6,9,12,15,18,21], (3, 7)),
            ("test 4", vec![50,100,150,200,250], (50, 5)),
            ("test 5", vec![100,200,300,400,500,600], (100, 6))
        ];

        for (desc, expect, input) in test_data {
            let got = count_by(input.0, input.1);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
