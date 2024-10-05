pub fn narcissistic(num: u64) -> bool {
    let length = num.to_string().len();
    let mut total = 0;
    let mut run_val = num;
    for _i in 0..length {
        let val = run_val%10;
        run_val = run_val / 10;
        total+= val.pow(length as u32);
    }

    total == num
}
