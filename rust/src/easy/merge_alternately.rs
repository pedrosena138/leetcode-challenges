struct Solution;
impl Solution {
    fn merge_alternately(word1: &String, word2: &String) -> String {
        let mut result = String::new();
        let w1_len = word1.len();
        let w2_len = word2.len();

        for i in 0..Self::max(w1_len, w2_len) {
            if i < w1_len {
                result.push(word1.to_lowercase().chars().nth(i).unwrap());
            }

            if i < w2_len {
                result.push(word2.to_lowercase().chars().nth(i).unwrap());
            }
        }
        result
    }
    fn max(a: usize, b: usize) -> usize {
        if a > b { a } else { b }
    }
}

#[test]
pub fn fist_test() {
    let word1 = &String::from("abc");
    let word2 = &String::from("pqr");
    let result = Solution::merge_alternately(&word1, &word2);
    assert_eq!(result, "apbqcr");
}

#[test]
pub fn second_test() {
    let word1 = &String::from("ab");
    let word2 = &String::from("pqrs");
    let result = Solution::merge_alternately(&word1, &word2);
    assert_eq!(result, "apbqrs");
}

#[test]
pub fn third_test() {
    let word1 = &String::from("abcd");
    let word2 = &String::from("pq");
    let result = Solution::merge_alternately(&word1, &word2);
    assert_eq!(result, "apbqcd");
}
