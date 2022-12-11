use std::char;

use std::collections::HashMap;

pub fn process_part1(input: &str) -> String {
    let letter_scores = ('a'..='z')
        .chain('A'..='Z')
        .enumerate()
        .map(|(idx, c)| (c, idx + 1))
        .collect::<HashMap<char, usize>>();

    let result = input
        .lines()
        .map(|line| {
            let compartment_size = line.len() / 2;
            let first_compartment = &line[0..compartment_size];
            let second_compartment = &line[compartment_size..];

            let common = first_compartment
                .chars()
                .find(|letter| second_compartment.contains(*letter))
                .unwrap();
            letter_scores.get(&common).unwrap()
        })
        .sum::<usize>();

    result.to_string()
}

pub fn process_part2(input: &str) -> String {
    let letter_scores = ('a'..='z')
        .chain('A'..='Z')
        .enumerate()
        .map(|(idx, c)| (c, idx + 1))
        .collect::<HashMap<char, usize>>();

    let result: Vec<_> = input.lines().collect();
    let mut common_letters: Vec<usize> = vec![];

    for i in 0..result.len()/3 {
        let common = result[3*i]
            .chars()
            .find(|letter| result[3*i + 1].contains(*letter) && result[3*i + 2].contains(*letter))
            .unwrap();
        
        common_letters.push(*letter_scores.get(&common).unwrap());
    }

    common_letters.iter().sum::<usize>().to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

    #[test]
    fn test1() {
        assert_eq!(process_part1(INPUT), "157");
    }

    #[test]
    fn test2() {
        assert_eq!(process_part2(INPUT), "70");
    }
}
