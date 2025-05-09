import numpy as np
from sklearn.datasets import fetch_openml
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, classification_report
from sklearn.model_selection import train_test_split

# Wczytanie zbioru MNIST
mnist = fetch_openml('mnist_784', version=1)
X = mnist.data
y = mnist.target.astype(np.uint8)

# Podział danych: 60000 do treningu, 10000 do testu
X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=10000, random_state=42, stratify=y
)

# Utworzenie i trenowanie klasyfikatora Random Forest
clf = RandomForestClassifier(n_estimators=100, random_state=42, n_jobs=-1)
clf.fit(X_train, y_train)

# Predykcja na zbiorze testowym
y_pred = clf.predict(X_test)

# Obliczenie dokładności (accuracy)
accuracy = accuracy_score(y_test, y_pred)
print("Dokładność (accuracy):", accuracy)

# Obliczenie czułości i precyzji dla poszczególnych klas
report = classification_report(y_test, y_pred, output_dict=True)
print("Czułość (macro avg):", report['macro avg']['recall'])
print("Precyzja (macro avg):", report['macro avg']['precision'])
