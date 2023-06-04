# Linked list implementation in python

class Node:

    def __init__(self, data, next=None):
        self.value = data
        self.next = next


class LinkedList:
    def __init__(self, NodeObj=None):
        self.head = NodeObj

    def add_node(self, data):
        new_node = Node(data=data)

        if self.head is None:
            self.head = new_node
        else:
            current = self.head
            while current.next is not None:
                current = current.next

            current.next = new_node

    def displayData(self):
        current = self.head
        while current is not None:
            print(f"{current.value} ->", end=' ')
            current = current.next
        print(end=' None')


if __name__ == '__main__':
    ll = LinkedList()
    for num in range(1, 11):
        ll.add_node(num)

    ll.displayData()
