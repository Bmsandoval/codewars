#[cfg(test)]
mod tests {
    use are_you_playing_banjo::are_you_playing_banjo;

    #[test]
    fn test_banjo() {
        let test_data = vec![
            ("test playing", "roger", "roger plays banjo"),
            ("test playing", "Ralph", "Ralph plays banjo"),
            ("test not playing", "jason", "jason does not play banjo"),
            ("test not playing", "Jack", "Jack does not play banjo")
        ];

        for (desc, input, expect) in test_data {
            let got = are_you_playing_banjo(input);
            assert_eq!(got, expect, "Failed on: {}", desc);
        }
    }
}
