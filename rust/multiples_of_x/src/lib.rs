pub fn count_by(x: u32, n: u32) -> Vec<u32> {
    // create vector to store results
    let mut resp = Vec::new();

    // loop n times
    for i in 0u32..n{
        // add value to vector x*current multiple
        resp.push(x*(i+1))
    }

    // return resp
    resp
}