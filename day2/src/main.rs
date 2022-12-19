use std::str::FromStr;

#[derive(Debug)]
enum Move {
    Rock,
    Paper,
    Scissors,
}

impl FromStr for Move {
    type Err = ();

    fn from_str(input: &str) -> Result<Move, Self::Err> {
        match input {
            "A" => Ok(Move::Rock),
            "B" => Ok(Move::Paper),
            "C" => Ok(Move::Scissors),
            _ => Err(()),
        }
    }
}

#[derive(Debug)]
enum Code {
    X,
    Y,
    Z,
}

impl FromStr for Code {
    type Err = ();

    fn from_str(input: &str) -> Result<Code, Self::Err> {
        match input {
            "X" => Ok(Code::X),
            "Y" => Ok(Code::Y),
            "Z" => Ok(Code::Z),
            _ => Err(()),
        }
    }
}

#[derive(Debug)]
enum Outcome {
    Lose,
    Draw,
    Win,
}

impl FromStr for Outcome {
    type Err = ();

    fn from_str(input: &str) -> Result<Outcome, Self::Err> {
        match input {
            "X" => Ok(Outcome::Lose),
            "Y" => Ok(Outcome::Draw),
            "Z" => Ok(Outcome::Win),
            _ => Err(()),
        }
    }
}

fn main() {
    let contents = include_str!("../input.txt");
    part_a(&contents);
    part_b(&contents);
}

fn part_a(contents: &str) {
    let mut total_points: u32 = 0;

    // Simply calculate the total of each elves
    for line in contents.lines() {
        let split = line.split(" ").collect::<Vec<&str>>();

        let opponent_move = Move::from_str(split[0]).unwrap();
        let my_move_code = Code::from_str(split[1]).unwrap();
        let my_move = get_move_from_code(my_move_code);
        let game_point = get_game_point(&opponent_move, &my_move);
        let my_move_point = get_move_point(&my_move);
        let total_point = game_point + my_move_point;
        total_points += total_point;
    }

    println!("Part A: {}", total_points);
}

// X means you need to lose,
// Y means you need to end the round in a draw,
// and Z means you need to win.
fn part_b(contents: &str) {
    let mut total_points: u32 = 0;

    // Simply calculate the total of each elves
    for line in contents.lines() {
        let split = line.split(" ").collect::<Vec<&str>>();

        let opponent_move = Move::from_str(split[0]).unwrap();
        let outcome = Outcome::from_str(split[1]).unwrap();
        let my_move = get_move_for_outcome(&opponent_move, &outcome);
        let game_point = get_game_point(&opponent_move, &my_move);
        let my_move_point = get_move_point(&my_move);
        let total_point = game_point + my_move_point;
        total_points += total_point;
    }

    println!("Part B: {}", total_points);
}

fn get_move_for_outcome(opponent_move: &Move, outcome: &Outcome) -> Move {
    match opponent_move {
        Move::Rock => {
            return match outcome {
                Outcome::Win => Move::Paper,
                Outcome::Lose => Move::Scissors,
                Outcome::Draw => Move::Rock,
            }
        }
        Move::Paper => {
            return match outcome {
                Outcome::Win => Move::Scissors,
                Outcome::Lose => Move::Rock,
                Outcome::Draw => Move::Paper,
            }
        }
        Move::Scissors => {
            return match outcome {
                Outcome::Win => Move::Rock,
                Outcome::Lose => Move::Paper,
                Outcome::Draw => Move::Scissors,
            }
        }
    }
}

fn get_move_from_code(m: Code) -> Move {
    match m {
        Code::X => Move::Rock,
        Code::Y => Move::Paper,
        Code::Z => Move::Scissors,
    }
}

fn get_game_point(opponent: &Move, us: &Move) -> u32 {
    match opponent {
        Move::Rock => match us {
            Move::Rock => 3,
            Move::Paper => 6,
            Move::Scissors => 0,
        },
        Move::Paper => match us {
            Move::Rock => 0,
            Move::Paper => 3,
            Move::Scissors => 6,
        },
        Move::Scissors => match us {
            Move::Rock => 6,
            Move::Paper => 0,
            Move::Scissors => 3,
        },
    }
}

fn get_move_point(my_move: &Move) -> u32 {
    match my_move {
        Move::Rock => 1,
        Move::Paper => 2,
        Move::Scissors => 3,
    }
}
