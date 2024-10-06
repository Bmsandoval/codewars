pub fn square_digits(num: u64) -> u64 {
    let digits = num.to_string().len();
    let mut resp = String::from("");
    let mut running_val = num;
    for _i in 0..digits {
        let cur_val = running_val%10;
        running_val/=10;
        resp = format!("{}{}",cur_val.pow(2), resp);
    }
    resp.parse::<u64>().unwrap()
}