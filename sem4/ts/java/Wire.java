import java.util.ArrayList;
import java.util.List;

public class Wire {
    private static final char DEFAULT_SIGNAL_SYMBOL = '_';
    private static final char OVERLAPPING_SIGNAL_SYMBOL = '#';
    private static final char JAM_SIGNAL_SYMBOL = '!';

    private final int length;
    private final char[] symbols;
    private final List<List<Signal>> signalGroups;
    private final List<SignalPropagator> signalPropagators;

    public int getLength() {
        return this.length;
    }
    
    public Wire(int length) {
        this.length = length;
        this.symbols = new char[length];
        this.signalGroups = new ArrayList<>();
        this.signalPropagators = new ArrayList<>();

        for (int i = 0; i < length; i++) {
            symbols[i] = DEFAULT_SIGNAL_SYMBOL;
            signalGroups.add(new ArrayList<>());
        }
    }

    public void tick() {
        List<SignalPropagator> activePropagators = new ArrayList<>();
        for (SignalPropagator propagator : signalPropagators) {
            propagator.tick();
            if (propagator.isActive()) {
                activePropagators.add(propagator);
            }
        }
        signalPropagators.clear();
        signalPropagators.addAll(activePropagators);

        for (int i = 0; i < length; i++) {
            List<Signal> activeSignals = new ArrayList<>();
            for (Signal signal : signalGroups.get(i)) {
                signal.tick();
                if (signal.isActive()) {
                    activeSignals.add(signal);
                }
            }
            signalGroups.set(i, activeSignals);
        }

        updateSegmentSymbols();
    }

    private void updateSegmentSymbols() {
        for (int i = 0; i < length; i++) {
            List<Signal> signals = signalGroups.get(i);
            if (signals.isEmpty()) {
                symbols[i] = DEFAULT_SIGNAL_SYMBOL;
            } else if (signals.size() == 1) {
                symbols[i] = signals.get(0).getSymbol();
            } else if (signals.stream().anyMatch(s -> s.getSymbol() == JAM_SIGNAL_SYMBOL)) {
                symbols[i] = JAM_SIGNAL_SYMBOL;
            } else {
                symbols[i] = OVERLAPPING_SIGNAL_SYMBOL;
            }
        }
    }

    public void addSignal(int devicePosition, char signalSymbol, int tickLifetime) {
        signalPropagators.add(new SignalPropagator(signalSymbol, devicePosition, tickLifetime, this));
    }

    public void addJamSignal(int devicePosition, int tickLifetime) {
        signalPropagators.add(new SignalPropagator(JAM_SIGNAL_SYMBOL, devicePosition, tickLifetime, this));
    }

    public boolean isFree(int position, char signalSymbol) {
        return symbols[position] == DEFAULT_SIGNAL_SYMBOL || symbols[position] == signalSymbol;
    }

    public boolean isCollision(int position, char signalSymbol) {
        return !isFree(position, signalSymbol);
    }

    public boolean isJammed(int position) {
        return symbols[position] == JAM_SIGNAL_SYMBOL;
    }

    @Override
    public String toString() {
        return new String(symbols);
    }

    private static class Signal {
        private final char symbol;
        private int ticksLeft;

        public Signal(char symbol, int ticksLeft) {
            this.symbol = symbol;
            this.ticksLeft = ticksLeft;
        }

        public void tick() {
            ticksLeft--;
        }

        public boolean isActive() {
            return ticksLeft > 0;
        }

        public char getSymbol() {
            return symbol;
        }
    }

    private static class SignalPropagator {
        private final char symbol;
        private int leftPosition;
        private int rightPosition;
        private final int tickLifetime;
        private int tickCounter;
        private final Wire wire;

        public SignalPropagator(char symbol, int position, int tickLifetime, Wire wire) {
            this.symbol = symbol;
            this.leftPosition = position;
            this.rightPosition = position;
            this.tickLifetime = tickLifetime;
            this.tickCounter = tickLifetime;
            this.wire = wire;
        }

        public void tick() {
            tickCounter--;
            if (leftPosition == rightPosition) {
                propagateSignal(leftPosition);
            } else {
                propagateSignal(leftPosition);
                propagateSignal(rightPosition);
            }
            leftPosition--;
            rightPosition++;
        }

        private void propagateSignal(int position) {
            if (isPositionValid(position)) {
                wire.signalGroups.get(position).add(new Signal(symbol, tickLifetime));
            }
        }

        private boolean isPositionValid(int position) {
            return position >= 0 && position < wire.length;
        }

        public boolean isActive() {
            return tickCounter > 0;
        }
    }
}
