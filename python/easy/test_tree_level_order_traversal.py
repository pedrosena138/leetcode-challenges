from collections import deque
import pytest


class Node:
    def __init__(self, info):
        self.info = info
        self.left = None
        self.right = None
        self.level = None

    def __str__(self):
        return str(self.info)


class BinarySearchTree:
    def __init__(self):
        self.root = None

    def create(self, val):
        if self.root == None:
            self.root = Node(val)
        else:
            current = self.root

            while True:
                if val < current.info:
                    if current.left:
                        current = current.left
                    else:
                        current.left = Node(val)
                        break
                elif val > current.info:
                    if current.right:
                        current = current.right
                    else:
                        current.right = Node(val)
                        break
                else:
                    break


"""
Node is defined as
self.left (the left child of the node)
self.right (the right child of the node)
self.info (the value of the node)
"""


def levelOrder(root):
    # Write your code here
    out = [root]
    queue = deque(out)
    while queue:
        for _ in range(len(queue)):
            node = queue.popleft()
            if node.left:
                queue.append(node.left)
                out.append(node.left)
            if node.right:
                queue.append(node.right)
                out.append(node.right)

    print(" ".join(map(lambda n: str(n.info), out)))


@pytest.mark.parametrize("t,arr", [(6, (1, 2, 5, 3, 6, 4))])
def test_case(t, arr):
    tree = BinarySearchTree()

    for i in range(t):
        tree.create(arr[i])

    levelOrder(tree.root)
