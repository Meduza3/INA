import numpy as np
import matplotlib.pyplot as plt

# Zakres wartości theta od 0 do 2*pi
theta = np.linspace(0, 2 * np.pi, 1000)
# Równanie krzywej czterolistnej
r = np.sin(2 * theta)

# Transformacja na współrzędne kartezjańskie
x = r * np.cos(theta)
y = r * np.sin(theta)

# Tworzenie wykresu
plt.figure(figsize=(6, 6))
plt.plot(x, y, label='r(θ) = sin(2θ)')
plt.xlabel('x')
plt.ylabel('y')
plt.title('Krzywa czterolistna')
plt.legend()
plt.axis('equal')  # Zapewnienie, że osie mają taką samą skalę
plt.grid(True)
plt.show()
