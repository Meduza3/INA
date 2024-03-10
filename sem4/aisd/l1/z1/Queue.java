public class Queue {
    private int[] queue;
    private int front;
    private int rear;
    private int capacity;
    private int count;

    public Queue(int size) {
        queue = new int[size];
        capacity = size;
        front = 0;
        rear = -1;
        count = 0;
    }

    public void enqueue(int item) {
        if (isFull()) {
            System.out.println("Queue is full");
            System.exit(1);
        }

        rear = (rear + 1) % capacity;
        queue[rear] = item;
        count++;
    }

    public int dequeue() {
        if (isEmpty()) {
            System.out.println("Queue is empty");
            System.exit(1);
        }

        int item = queue[front];
        front = (front + 1) % capacity;
        count--;
        return item;
    }
    
    public int size() {
        return count;
    }

    public boolean isEmpty() {
        return (size() == 0);
    }

    public boolean isFull() {
        return (size() == capacity);
    }
}