import subprocess

def main():
    
    for n in range(1000, 50000, 1000):
        subprocess.run(f'./generator_losowy {n} | ./sort -insertion', shell=True)
        subprocess.run(f'./generator_losowy {n} | ./sort -quick', shell=True)
        subprocess.run(f'./generator_losowy {n} | ./sort -hybrid', shell=True)
        subprocess.run(f'./generator_losowy {n} | ./sort -merge', shell=True)
        subprocess.run(f'./generator_losowy {n} | ./sort -dualpivot', shell=True)
        subprocess.run(f'./generator_losowy {n} | ./sort -custom', shell=True)


if __name__ == "__main__":
    main()