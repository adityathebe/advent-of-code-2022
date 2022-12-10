fn main() {
    let contents = include_str!("../input.txt");

    let mut current: i32 = 0;
    let mut total_calories: Vec<i32> = Vec::new();

    // Simply calculate the total of each elves
    for line in contents.lines() {
        if line == "" {
            total_calories.push(current);
            current = 0;
            continue;
        }

        current += line.parse::<i32>().unwrap();
    }

    total_calories.push(current);
    total_calories.sort();
    total_calories.reverse();

    let sum_of_three: i32 = total_calories.as_slice()[..3].iter().sum();
    println!("Sum of top 3: {}", sum_of_three);
}

fn _part_a() {
    let contents = include_str!("../input.txt");

    // current is the total calories counted so far
    // by a given Elf.
    let mut current: i32 = 0;

    // max is the max number of calories counted by
    // one of Elves visted so far.
    let mut max: i32 = -1;

    for line in contents.lines() {
        if line == "" {
            if current > max {
                max = current
            }

            current = 0;
            continue;
        }

        current += line.parse::<i32>().unwrap();
    }

    println!("max {max}")
}
