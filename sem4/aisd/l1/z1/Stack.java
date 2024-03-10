public class Stack {
    private int[] stack;
    private int top;
    private int capacity;

    public Stack(int size) {
        stack = new int[size];
        capacity = size;
        top = -1;
    }

    public void push(int item) {
        if(isFull()) {
            System.out.println("Stack is full");
            System.exit(1);
        }
        stack[++top] = item;
    }

    public int pop() {
        if (isEmpty()){
            System.out.println("Stack is empty");
            System.exit(1);
        }
        return stack[top--];
    }

    public int size() {
        return top + 1;
    }

    public boolean isEmpty() {
        return top == -1;
    }

    public boolean isFull() {
        return top == capacity - 1;
    }
}


