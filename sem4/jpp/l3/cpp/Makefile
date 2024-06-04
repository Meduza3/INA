# Compiler and flags
CC = g++
FLAGS = -g -c -Wall
LFLAGS = 

# Source files and object files
OBJS = DHSetup.o Gf_int.o User.o main.o
SOURCE = DHSetup.cpp Gf_int.cpp User.cpp main.cpp
HEADER = DHSetup.hpp Gf_int.hpp User.hpp
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

# Clean rule
clean:
	rm -f $(OBJS) $(OUT)
