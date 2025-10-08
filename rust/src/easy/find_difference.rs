use std::collections::HashMap;
pub struct Solution;
impl Solution {
    pub fn find_difference(s: String, t: String) -> char {
        let mut char_counter: HashMap<char, i32> = HashMap::new();
        for c in s.chars() {
            *char_counter.entry(c).or_insert(0) -= 1;
        }
        for c in t.chars() {
            *char_counter.entry(c).or_insert(0) += 1;
        }

        let result: &char = char_counter
            .iter()
            .filter(|x| *x.1 == 1)
            .collect::<Vec<_>>()[0]
            .0;

        *result
    }
}

#[test]
pub fn fist_test() {
    let s = String::from("abcd");
    let t = String::from("abcde");
    let result = Solution::find_difference(s, t);
    assert!(result.is_ascii_alphabetic());
    assert_eq!(result, 'e');
}

#[test]
pub fn second_test() {
    let s = String::from("");
    let t = String::from("y");
    let result = Solution::find_difference(s, t);
    assert!(result.is_ascii_alphabetic());
    assert_eq!(result, 'y');
}

#[test]
pub fn third_test() {
    let s = String::from("a");
    let t = String::from("aa");
    let result = Solution::find_difference(s, t);
    assert_eq!(result, 'a');
}
