import matplotlib.pyplot as plt
import pandas as pd

file_path = 'output.csv'
data = pd.read_csv(file_path)

# Set up the plotting environment
plt.figure(figsize=(10, 6))

# Plot n vs timePrim and n vs timeKruskal on the same graph
plt.plot(data['n'], data['timePrim'], label='Prim', color='blue', marker='o')
plt.plot(data['n'], data['timeKruskal'], label='Kruskal', color='green', marker='o')

plt.title('n vs Time for Prim and Kruskal Algorithms')
plt.xlabel('Number of vertices (n)')
plt.ylabel('Time in milliseconds')
plt.grid(True)
plt.legend()

plt.show()
