pub fn process_part1(input: &str) -> String {
    let result = input
        .split("\n")
        .map(|line| {
            match line {
                "A X" => 4,
                "A Y" => 8,
                "A Z" => 3,
                "B X" => 1,
                "B Y" => 5,
                "B Z" => 9,
                "C X" => 7,
                "C Y" => 2,
                "C Z" => 6,
                _ => 0
            }
        })
        .sum::<u32>();
    result.to_string()
}

pub fn process_part2(input: &str) -> String {
    let result = input
        .split("\n")
        .map(|line| {
            match line {
                "A X" => 3,
                "A Y" => 4,
                "A Z" => 8,
                "B X" => 1,
                "B Y" => 5,
                "B Z" => 9,
                "C X" => 2,
                "C Y" => 6,
                "C Z" => 7,
                _ => 0
            }
        })
        .sum::<u32>();
    result.to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "A X
B Y
C Z";

    #[test]
    fn test1() {
        let result = process_part1(INPUT);
        assert_eq!(result, "15");
    }

    #[test]
    fn test2() {
        let result = process_part2(INPUT);
        assert_eq!(result, "15");
    }
}
