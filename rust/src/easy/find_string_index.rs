pub struct Solution;
impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        let mut result: i32 = -1;
        let needle_size = needle.len();
        let haystack_size = haystack.len();
        let mut sliced_haystack: Vec<String> = [].to_vec();

        let mut i = 0;
        while i < haystack_size {
            let pivot = i + needle_size;
            let mut sub_str: &str= "";
            if pivot > haystack_size {
                sub_str = &haystack[i..haystack_size];
            } else {
                sub_str = &haystack[i..i + needle_size];
            }
            sliced_haystack.push(sub_str.to_string());
            i += 1;
        }

        for (i, n) in sliced_haystack.iter().enumerate() {
            if n == &needle {
                result = i as i32;
                return result
            }
        } 

        result
    }
}

#[test]
pub fn fist_test() {
    let haystack = String::from("sadbutsad");
    let needle = String::from("sad");
    let result = Solution::str_str(haystack, needle);
    assert_eq!(result, 0);
}

#[test]
pub fn second_test() {
    let haystack = String::from("leetcode");
    let needle = String::from("leeto");
    let result = Solution::str_str(haystack, needle);
    assert_eq!(result, -1);
}

#[test]
pub fn third_test() {
    let haystack = String::from("hello");
    let needle = String::from("ll");
    let result = Solution::str_str(haystack, needle);
    assert_eq!(result, 2);
}