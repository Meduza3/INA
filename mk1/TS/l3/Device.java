import java.util.EnumMap;
import java.util.Map;
import java.util.Random;

enum State {
    RECEIVING,
    TRANSMITTING,
    WAITING,
    JAMMING,
    BACKOFF
}

public class Device {

    private static final int MAX_BACKOFF_TRANSMISSION_ATTEMPTS = 16;
    public Map<State, Integer> timeInStates = new EnumMap<>(State.class);
    public int totalCollisions = 0;

    private final String symbol;
    private final Wire wire;
    private final int positionInWire;

    private final int minPacketTime;
    private int tickCounter;
    private int retransmissionsCounter;
    private int backoffWaitingTicks;
    private State state;
    private boolean isReadyToTransmit;

    private int successfulTransmissions;
    private int failedTransmissions;

    public Device(String symbol, Wire wire, int positionInWire) {
        this.symbol = symbol;
        this.wire = wire;
        this.positionInWire = positionInWire;

        this.minPacketTime = 2 * wire.getLength();
        this.tickCounter = 0;
        this.retransmissionsCounter = 0;
        this.backoffWaitingTicks = 0;
        this.state = State.RECEIVING;
        this.isReadyToTransmit = true;

        this.successfulTransmissions = 0;
        this.failedTransmissions = 0;

        for (State state : State.values()) {
            timeInStates.put(state, 0); // Initialize time spent in each state to zero
        }
    }

    public int getSuccessfulTransmissions() {
        return successfulTransmissions;
    }

    public int getFailedTransmissions() {
        return failedTransmissions;
    }

    public void tick() {
        timeInStates.put(state, timeInStates.get(state) + 1); // Increment state time
        switch (state) {
            case TRANSMITTING:
                transmit();
                break;
            case WAITING:
                sendPacket();
                break;
            case JAMMING:
                jam();
                break;
            case BACKOFF:
                backoff();
                break;
        }
    }

    public void sendPacket() {
        if (isReadyToTransmit) {
            if (wire.isCollision(positionInWire, symbol.charAt(0))) {
                totalCollisions++; // Increment collision count
                state = State.WAITING;
            } else {
                state = State.TRANSMITTING;
                tickCounter = 0;
                wire.addSignal(positionInWire, symbol.charAt(0), minPacketTime);
                isReadyToTransmit = false;
            }
        }
    }

    private void transmit() {
        if (wire.isCollision(positionInWire, symbol.charAt(0))) {
            if (wire.isJammed(positionInWire)) {
                state = State.BACKOFF;
                retransmissionsCounter = 0;
            } else {
                state = State.JAMMING;
                wire.addJamSignal(positionInWire, minPacketTime);
                tickCounter = 0;
            }
            failedTransmissions++;
        } else {
            tickCounter++;
            if (tickCounter == minPacketTime) {
                state = State.RECEIVING;
                isReadyToTransmit = true;
                successfulTransmissions++;
            }
        }
    }

    private void jam() {
        tickCounter++;
        if (tickCounter == minPacketTime) {
            state = State.BACKOFF;
            retransmissionsCounter = 0;
            calculateExponentialBackoff();
        }
    }

    private void backoff() {
        backoffWaitingTicks--;
        if (backoffWaitingTicks <= 0) {
            if (wire.isFree(positionInWire, symbol.charAt(0))) {
                state = State.TRANSMITTING;
                tickCounter = 0;
                wire.addSignal(positionInWire, symbol.charAt(0), minPacketTime);
            } else {
                calculateExponentialBackoff();
                failedTransmissions++;
            }
        }
    }

    private void calculateExponentialBackoff() {
        if (retransmissionsCounter == MAX_BACKOFF_TRANSMISSION_ATTEMPTS) {
            throw new RuntimeException("Max transmission attempts reached.");
            }
            Random random = new Random();
            int k = Math.min(retransmissionsCounter, 10);
        retransmissionsCounter++;
        backoffWaitingTicks = random.nextInt((1 << k)) * minPacketTime; // random from 0 to 2^k
        // Moze byc 0 dzieki temu ze k moze byc rowne 0 wiec wtedy losuje miedzy 0 i 1
    }

    public String getSymbol() {
        return this.symbol;
    }

    public Map<State, Integer> getTimeInStates() {
            return timeInStates;

    }
}
