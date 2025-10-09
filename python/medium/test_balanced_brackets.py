#!/bin/python3
import pytest

#
# Complete the 'isBalanced' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

TOKENS = {"(": ")", "[": "]", "{": "}"}


def isBalanced(s):
    stack = []

    for i in s:
        if i in TOKENS.keys():
            stack.append(i)
        else:
            if stack:
                if TOKENS[stack.pop()] != i:
                    return "NO"
            else:
                return "NO"
    return "NO" if stack else "YES"


@pytest.mark.parametrize(
    "s, expected", [("{[()]}", "YES"), ("{[(])}", "NO"), ("{{[[(())]]}}", "YES")]
)
def test_case_0(s, expected):
    result = isBalanced(s)
    assert result == expected


@pytest.mark.parametrize("s, expected", [("{{([])}}", "YES"), ("{{)[](}}", "NO")])
def test_case_1(s, expected):
    result = isBalanced(s)
    assert result == expected


@pytest.mark.parametrize(
    "s, expected",
    [("{(([])[])[]}", "YES"), ("{(([])[)[]]}", "NO"), ("{(([])[])[]}[]", "YES")],
)
def test_case_2(s, expected):
    result = isBalanced(s)
    assert result == expected
