#!/usr/bin/env python3

import matplotlib.pyplot as plt

x_values = []
y_values = []

# Read data from the file
with open('select-7.txt', 'r') as file:
    for line in file:
        x, y, z = line.split()  # Assuming the values are separated by whitespace
        x_values.append(float(x))
        y_values.append(float(z))

# Calculate the average of the corresponding y values
averages = {}
for x, y in zip(x_values, y_values):
    if x not in averages:
        averages[x] = []
    averages[x].append(y)

x_avg = []
y_avg = []
for x, y_list in averages.items():
    x_avg.append(x)
    y_avg.append(sum(y_list) / len(y_list))

# Plot the averages
plt.plot(x_avg, y_avg)
plt.xlabel('Rozmiar tablicy')
plt.ylabel('Średnia ilość przestawień')
plt.title('Przestawienia - Select-7')
plt.savefig('swaps_select7.png')