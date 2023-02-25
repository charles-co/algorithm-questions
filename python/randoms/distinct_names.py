# You are given an array of strings ideas that represents a list of names to be used in the process of naming a company. The process of naming a company is as follows:

# Choose 2 distinct names from ideas, call them ideaA and ideaB.
# Swap the first letters of ideaA and ideaB with each other.
# If both of the new names are not found in the original ideas, then the name ideaA ideaB (the concatenation of ideaA and ideaB, separated by a space) is a valid company name.
# Otherwise, it is not a valid name.
# Return the number of distinct valid names for the company.


from typing import List

class Solution:
    def distinctNames(self, ideas: List[str]) -> int:
        groups = [set() for _ in range(26)]
        for idea in ideas:
            groups[ord(idea[0]) - ord("a")].add(idea[1:])

        result = 0

        for i in range(0, 25):
            for j in range(i + 1, 26):
                mutuals = len(groups[i] & groups[j])
                result += 2 * (len(groups[i]) - mutuals) * (len(groups[j]) - mutuals)
        return result
    