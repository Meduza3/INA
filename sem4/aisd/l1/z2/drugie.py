class Node:
    def __init__(self, value=None):
        self.value = value
        self.next = None


class LinkedList:
    def __init__(self):
        self.head = None
        self.size = 0

    def insert(self, value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            new_node.next = self.head
        else:
            current = self.head
            while current.next != self.head:
                current = current.next
            current.next = new_node
            new_node.next = self.head
        self.size += 1

    def merge(self, other):
        if self.head is None:
            self.head = other.head
            return
        if other.Head is None:
            return
        current = self.head
        while current.next != self.head:
            current = current.next
        current = other.head
        while current.next != other.head:
            current = current.next
        current.next = self.head
        self.size += other.size

    def display(self):
        elements = []
        current = self.head
        if self.head is not None:
            while True:
                elements.append(current.value)
                current = current.next
                if current == self.head:
                    break
        return elements

import random

T = [random.randint(0, 100000) for _ in range(10000)]
L = LinkedList()
for number in T:
    L.insert(number)

def search_cost(l, search_values):
    total_comparisons = 0
    for value in search_values:
        current = l.head
        comparisons = 0
        while True:
            comparisons += 1
            if current.value == value:
                break
            current = current.next
            if current == l.head:
                break
        total_comparisons += comparisons
    return total_comparisons / len(search_values)

search_values_existing = random.sample(T, 1000)
cost_existing = search_cost(L, search_values_existing)

search_values_random = [random.randint(0, 100000) for _ in range(1000)]
cost_random = search_cost(L, search_values_random)

print(cost_existing)
print(cost_random)
