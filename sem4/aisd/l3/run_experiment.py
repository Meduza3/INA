#!/usr/bin/env python3


import subprocess

for n in range(0, 50100, 100):
    for _ in range(50):
        command = f"./generator {n} | ./select-aisd -select-9"
        subprocess.run(command, shell=True)
