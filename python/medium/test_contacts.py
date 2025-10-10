class Node:
    def __init__(self):
        self.children = {}
        self.size = 0

    def __repr__(self):
        return f"Node<size={self.size}>"


class Trie:
    root = Node()

    def add(self, name: str):
        node = self.root

        for c in name:
            if c not in node.children:
                node.children[c] = Node()

            node = node.children[c]
            node.size += 1

    def find(self, name):
        node = self.root

        for c in name:
            if c not in node.children:
                return 0
            node = node.children[c]

        return node.size


def contacts(queries):
    trie = Trie()
    result = []

    for op, name in queries:
        if op == "add":
            trie.add(name)
        if op == "find":
            result.append(trie.find(name))

    return result
