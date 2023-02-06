# Trie is a tree data structure that is used to store strings. It is used to store a dictionary of words.
import unittest
from collections import Counter


class TrieNode:
    def __init__(self):
        self.children = {}
        self.end_of_word = False


class Trie:
    def __init__(self):
        self.root = TrieNode()

    def insert(self, word):
        current = self.root
        for char in word:
            if char not in current.children:
                current.children[char] = TrieNode()
            current = current.children[char]
        current.end_of_word = True

    def search(self, word, char_counts):
        current = self.root
        for char in word:
            if char not in current.children:
                return False
            if char_counts[char] == 0:
                return False
            char_counts[char] -= 1
            current = current.children[char]
        return current.end_of_word


def longest_matching_word(dictionary, chars):
    trie = Trie()
    char_counts = Counter(chars)

    for word in dictionary:
        trie.insert(word)

    result = ""
    for word in dictionary:
        count = char_counts.copy()
        if trie.search(word, count):
            if len(word) > len(result):
                result = word
    return result
        

print(longest_matching_word(['when', 'what', 'whatthen', 'whatnow'], "whatno"))
print(longest_matching_word(['when', 'what', 'whatthen', 'whatnow'], "whatnwo"))
print(longest_matching_word(['when', 'what', 'whatthen', 'whatnow'], "wontthaw"))
