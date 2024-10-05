pub fn are_you_playing_banjo(name: &str) -> String {
    if name.to_lowercase().starts_with('r') {
        return (name.to_owned() + " plays banjo").to_string()
    }

    (name.to_owned() + " does not play banjo").to_string()
}