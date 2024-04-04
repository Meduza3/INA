import numpy as np
import matplotlib.pyplot as plt

# Definicja zakresu i siatki
x = np.linspace(-5, 5, 100)
y = np.linspace(-5, 5, 100)
x, y = np.meshgrid(x, y)

# Tworzenie figure
fig = plt.figure(figsize=(20, 10))

# 1. V(z - x^2 - y^2)
ax1 = fig.add_subplot(221, projection='3d')
z = x**2 + y**2
ax1.plot_surface(x, y, z, cmap='viridis', alpha=0.6)
ax1.set_title('V(z - x^2 - y^2)')

# 2. V(z^2 - x^2 - y^2)
ax2 = fig.add_subplot(222, projection='3d')
z = np.sqrt(x**2 + y**2)
ax2.plot_surface(x, y, z, cmap='viridis', alpha=0.6)
ax2.plot_surface(x, y, -z, cmap='viridis', alpha=0.6)
ax2.set_title('V(z^2 - x^2 - y^2)')

# 3. V(z - x^2 + y^2)
ax3 = fig.add_subplot(223, projection='3d')
z = x**2 - y**2
ax3.plot_surface(x, y, z, cmap='viridis', alpha=0.6)
ax3.set_title('V(z - x^2 + y^2)')

# 4. V(xz, yz)
ax4 = fig.add_subplot(224, projection='3d')
z = np.linspace(-5, 5, 100)
X, Z = np.meshgrid(x, z)
Y, Z = np.meshgrid(y, z)
# Dla z = 0
ax4.plot_surface(X, Y*0, Z, cmap='viridis', alpha=0.6)
# Osi x i y
ax4.plot(np.zeros(100), np.zeros(100), z, color='black')
ax4.set_title('V(xz, yz)')

plt.show()
