# Compiler and flags
CC = g++
FLAGS = -g -c -Wall
LFLAGS = 

# Source files and object files
OBJS = gcd.o main.o test.o
SOURCE = gcd.cpp main.cpp test.cpp
HEADER = gcd.h test.h
OUT = main

# Libraries
LDLIBS = 

# Default target
all: $(OUT)

# Linking rules
$(OUT): $(OBJS)
	$(CC) -g $(OBJS) -o $(OUT) $(LFLAGS) $(LDLIBS)

# Compilation rules
%.o: %.cpp $(HEADER)
	$(CC) $(FLAGS) -o $@ $<

#test rule

test: all
	./$(OUT)

	
# Clean rule
clean:
	rm -f $(OBJS) $(OUT)
