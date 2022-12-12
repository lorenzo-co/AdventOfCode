pub fn process_part1(input: &str) -> String {
    let result = input
        .lines()
        .map(|line| 
            line
                .split(",")
                .map(|sectors| sectors.split("-"))
        )
        .map(|line| line.flatten())
        .fold(0, |acc, line| {
            let values: Vec<i32> = line.clone().map(|x| x.parse::<i32>().unwrap()).collect();

            if (
                values[0] <= values[2] 
                && values[1] >= values[3] 
                && values[0] <= values[3] 
                && values[1] >= values[2]
            ) || (
                values[0] >= values[2] 
                && values[1] <= values[3] 
                && values[0] <= values[3]
                && values[1] >= values[2]
            ) {
                return acc + 1
            } else {
                return acc
            }
        });
    
    result.to_string()
}

pub fn process_part2(input: &str) -> String {
    let result = input
        .lines()
        .map(|line| 
            line
                .split(",")
                .map(|sectors| sectors.split("-"))
        )
        .map(|line| line.flatten())
        .fold(0, |acc, line| {
            let values: Vec<i32> = line.clone().map(|x| x.parse::<i32>().unwrap()).collect();

            if (
                values[0] <= values[2] 
                && values[1] >= values[2] 
            ) || (
                values[0] >= values[2] 
                && values[0] <= values[3]
            ) {
                return acc + 1
            } else {
                return acc
            }
        });
    
    result.to_string()
}

#[cfg(test)]
mod tests {
    use crate::{process_part1, process_part2};

    const INPUT: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

    #[test]
    fn test1() {
        assert_eq!(process_part1(INPUT), "2");
    }

    #[test]
    fn test2() {
        assert_eq!(process_part2(INPUT), "4");
    }
}
