import subprocess

def main():
    
    for n in range(1000, 50000, 1000):
        subprocess.run(f'python3 generator_losowy.py {n} | ./sorter -insertion', shell=True)
        subprocess.run(f'python3 generator_losowy.py {n} | ./sorter -quick', shell=True)
        subprocess.run(f'python3 generator_losowy.py {n} | ./sorter -hybrid', shell=True)
        subprocess.run(f'python3 generator_losowy.py {n} | ./sorter -merge', shell=True)
        subprocess.run(f'python3 generator_losowy.py {n} | ./sorter -dualpivot', shell=True)
        subprocess.run(f'python3 generator_losowy.py {n} | ./sorter -custom', shell=True)


if __name__ == "__main__":
    main()