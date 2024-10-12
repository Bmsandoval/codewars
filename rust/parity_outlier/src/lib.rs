pub fn find_outlier(values: &[i32]) -> i32 {
    let mut resp = 0i32;
    let mut first: i32 = 0;
    let mut second: i32 = 0;
    let mut third = false;
    let mut is_even = false;
    for val in values {
        if *val == 0 {
            return 0
        }
        if first == 0 {
            first = *val;
            continue
        } else if second == 0 {
            second = *val;
            continue
        } else if third == false {
            third = true;
            let mut odd_count = 0;
            let mut even_count = 0;
            if first % 2 > 0 {
                odd_count +=1;
            } else {
                even_count +=1
            }
            if second % 2 > 0 {
                odd_count +=1;
            } else {
                even_count +=1
            }
            if *val % 2 > 0 {
                odd_count +=1;
            } else {
                even_count +=1
            }
            if even_count > odd_count {
                is_even = true;
            }

            if is_even && odd_count > 0 {
                if first % 2 > 0 {
                    resp = first;
                    break;
                }
                if second % 2 > 0 {
                    resp = second;
                    break;
                }
                if *val % 2 > 0 {
                    resp = *val;
                    break;
                }
            } else if !is_even && even_count > 0 {
                if first % 2 < 1 {
                    resp = first;
                    break;
                }
                if second % 2 < 1 {
                    resp = second;
                    break;
                }
                if *val % 2 < 1 {
                    resp = *val;
                    break;
                }
            }
        }
        if (is_even && *val % 2 > 0) || (!is_even && *val % 2 < 1) {
            resp = *val;
        }
    }
    resp
}