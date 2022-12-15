use std::collections::BTreeSet;

pub fn process_part1(input: &str) -> String {
    const WIN_SIZE: usize = 4;

    let letters = input
        .chars()
        .collect::<Vec<char>>();
    let result = letters.windows(WIN_SIZE)
        .enumerate()
        .find(|(_i, four_letters)| {
            let unique_letters =
                four_letters.iter().collect::<BTreeSet<&char>>();
            four_letters.len() == unique_letters.len()
        }).unwrap();
    let (index, _) = result;
    (index+WIN_SIZE).to_string()
}

pub fn process_part2(input: &str) -> String {
    const WIN_SIZE: usize = 14;
    let letters = input
        .chars()
        .collect::<Vec<char>>();
    let result = letters.windows(WIN_SIZE)
        .enumerate()
        .find(|(_i, four_letters)| {
            let unique_letters =
                four_letters.iter().collect::<BTreeSet<&char>>();
            four_letters.len() == unique_letters.len()
        }).unwrap();
    let (index, _) = result;
    (index+WIN_SIZE).to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test1() {
        let result = process_part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb");
        assert_eq!(result, "7");
    }

    #[test]
    fn test2() {
        let result = process_part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb");
        assert_eq!(result, "19");
    }
}
