import random

class Node:
    def __init__(self, value=None):
        self.value = value
        self.next = None
        self.prev = None

class LinkedList:
    def __init__(self):
        self.head = None
        self.size = 0

    def insert(self, value):
        new_node = Node(value)
        if self.head is None:
            self.head = new_node
            new_node.next = new_node.prev = new_node
        else:
            tail = self.head.prev
            tail.next = new_node
            new_node.prev = tail
            new_node.next = self.head
            self.head.prev = new_node
        self.size += 1
    
    def merge(self, other):
        if self.head is None:
            self.head = other.head
            return
        if other.head is None:
            return
        self_tail = self.head.prev
        other_tail = other.head.prev
        self_tail.next = other.head
        other.head.prev = self_tail
        other_tail.next = self.head
        self.head.prev = other_tail
        self.size += other.size

    def display(self):
        elements = []
        current = self.head
        if self.head is not None:
            for _ in range(self.size):
                elements.append(current.value)
                current = current.next
        return elements

def search_cost(l, search_values):
    total_comparisons = 0
    for value in search_values:
        current = l.head
        comparisons = 0
        direction_forward = random.choice([True, False])
        while True:
            comparisons += 1
            if current.value == value:
                break
            current = current.next if direction_forward else current.prev
            if comparisons >= l.size:
                break
        total_comparisons += comparisons
    return total_comparisons / len(search_values)

T = [random.randint(0, 100000) for _ in range(10000)]
L = LinkedList()

for number in T:
    L.insert(number)



search_values_existing = random.sample(T, 1000)
search_values_random = [random.randint(0, 100000) for _ in range(1000)]

cost_existing = search_cost(L, search_values_existing)
cost_random = search_cost(L, search_values_random)

print(cost_existing)
print(cost_random)