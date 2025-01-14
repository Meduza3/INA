import matplotlib.pyplot as plt

# Sample data
data = {
    "DaneXXX": [16, 10000, 100000, 300000, 500000],
    "no_pivot": [1.971412604, 2.025456928, 2.102957229, 2.439017346, 2.741120299],
    "pivot": [0.939010455, 0.958661113, 1.011004015, 1.73092819, 1.537840156],
    "standard": [2.04367760, 2.305941296, 2.204091731, 3.113145581, 2.940711207]
}

# Extracting values for plotting
x = data["DaneXXX"]
y_pivot = data["pivot"]
y_no_pivot = data["no_pivot"]
y_standard = data["standard"]

# Plotting the lines with logarithmic scale
plt.plot(x, y_pivot, label='Z wyborem elementu głównego', marker='o')
plt.plot(x, y_no_pivot, label='Bez wyboru elementu', marker='s')
plt.plot(x, y_standard, label='Metoda tradycyjna', marker='^')

# Adding labels and title
plt.xlabel('Rozmiar macierzy (Skala logarytmiczna)')
plt.ylabel('Sekundy')
plt.title('Porównanie czasu wykonywania Eliminacji')

# Setting logarithmic scale for x-axis
plt.xscale('log')

# Adding legend
plt.legend()

# Display the plot
plt.grid(True, which="both", linestyle="--", linewidth=0.5)
plt.savefig("wykres.png")
