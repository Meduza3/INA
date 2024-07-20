import matplotlib.pyplot as plt
import pandas as pd

file_path = 'output.csv'
data = pd.read_csv(file_path)

# Set up the plotting environment
plt.figure(figsize=(10, 6))

# Plot n vs timePrim and n vs timeKruskal on the same graph
plt.plot(data['Number of Nodes'], data['Average Rounds'], label='Average Rounds', color='blue', marker='o')
plt.plot(data['Number of Nodes'], data['Minimum Rounds'], label='Minimum Rounds', color='green', marker='o')
plt.plot(data['Number of Nodes'], data['Maximum Rounds'], label='Maximum Rounds', color='red', marker='o')

plt.title('MST Graph')
plt.xlabel('Number of vertices (n)')
plt.ylabel('Rounds needed')
plt.grid(True)
plt.legend()

plt.savefig('plot2.png')
