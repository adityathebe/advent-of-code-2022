fn main() {
    let contents = include_str!("../input.txt");
    let lines = contents.lines();

    let mut current: i32 = 0;
    let mut max: i32 = -1;
    for line in lines {
        if line == "" {
            if current > max {
                max = current
            }

            current = 0;
            continue;
        }

        let res = line.parse::<i32>();
        match res {
            Ok(x) => current += x,
            Err(e) => println!("{e}"),
        }
    }

    println!("max {max}")
}
