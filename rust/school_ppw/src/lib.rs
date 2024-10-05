pub fn paperwork(n: i16, m: i16) -> u32 {
    if n < 1 || m < 1 {
        0u32
    } else {
        (n * m) as u32
    }
}