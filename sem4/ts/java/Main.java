public class Main {
    public static void main(String[] args) {
        final int WIRE_LENGTH = 30;
        final double TRANSMISSION_PROBABILITY = 0.005;
        final int TICKS = 2000;

        Wire wire = new Wire(WIRE_LENGTH);
        Device[] devices = {
            new Device("A", wire, 5),
            new Device("B", wire, 15),
            new Device("C", wire, 25)
        };

        for (int i = 0; i < TICKS; i++) {
            for (Device device : devices) {
                if (Math.random() < TRANSMISSION_PROBABILITY) {
                    device.sendPacket();
                }
                device.tick();
            }
            wire.tick();

            // Clear the current line and print the new state of the wire
            System.out.print("\r" + wire);  // Carriage return to rewrite the line in the console
            try {
                Thread.sleep(10);  // Sleep for a short time (e.g., 100ms) to make the output human-readable
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();  // Handle thread interruption
                System.err.println("Thread was interrupted, failed to complete operation");
            }
        }

        System.out.println();  // Move to the next line after the loop completes
        for (Device device : devices) {
            System.out.printf("Device %s: %d successful transmissions, %d failed transmissions.\n",
                    device.getSymbol(), device.getSuccessfulTransmissions(), device.getFailedTransmissions());
                    for (State state : State.values()) {
                        Integer time = device.getTimeInStates().getOrDefault(state, 0);
                        System.out.printf("    Time in %s: %d ticks\n", state, time);
                    }
        }
    }
}
