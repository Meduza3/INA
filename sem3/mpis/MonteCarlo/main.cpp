#include <math.h>
#include <iostream>
#include <random>
#include <fstream>

double calculateIntegral(double (*f)(double), double a, double b, int n) {

    double h = (b-a)/1000;
    double x = a;
    double M = f(x);
    double C = 0;

    while (x < b) {
        x += h;
        if (f(x) > M) {
            M = f(x);
        }
    }

    
    std::mt19937 mt(std::random_device{}());
    std::uniform_real_distribution<double> dist_x(a, b);
    std::uniform_real_distribution<double> dist_y(0, M);

    for(int i = 0; i < n; i++) {
        double x = dist_x(mt);
        double y = dist_y(mt);
        if (y <= f(x)) {
            C++;
        }
    }

    return C/n * (b-a) * M;
}

double f(double x) {
    return pow(x, 1.0/3.0);
}

double g(double x) {
    return sin(x);
}

double h(double x) {
    return 4*x*pow(1-x,3);
}

double half_pi(double x) {
    return 1/(pow(1-x*x, 0.5));
}

int main() {


    std::ofstream outfile_f("data/data_f.txt");
    std::ofstream outfile_g("data/data_g.txt");
    std::ofstream outfile_h("data/data_h.txt");
    std::ofstream outfile_pi("data/data_pi.txt");

    std::cout.rdbuf(outfile_f.rdbuf());
    for(int n = 50; n <= 5000; n += 50) {
        double avg = 0;
        for(int i = 0; i < 50; i++) {
            double res = calculateIntegral(f, 0, 8, n);
            std::cout << n << " " << res; 
            avg += res;
            if(i == 49){
                std::cout << " " << avg/50;
            }
            std::cout << std::endl;
        }
    }

    std::cout.rdbuf(outfile_g.rdbuf());
    for(int n = 50; n <= 5000; n += 50) {
        double avg = 0;
        for(int i = 0; i < 50; i++) {
            double res = calculateIntegral(g, 0, M_PI, n);
            std::cout << n << " " << res; 
            avg += res;
            if(i == 49){
                std::cout << " " << avg/50;
            }
            std::cout << std::endl;
        }
    }

    std::cout.rdbuf(outfile_h.rdbuf());
    for(int n = 50; n <= 5000; n += 50) {
        double avg = 0;
        for(int i = 0; i < 50; i++) {
            double res = calculateIntegral(h, 0, 1, n);
            std::cout << n << " " << res; 
            avg += res;
            if(i == 49){
                std::cout << " " << avg/50;
            }
            std::cout << std::endl;
        }
    }

    std::cout.rdbuf(outfile_pi.rdbuf());
    for(int n = 50; n <= 5000; n += 50) {
        double avg = 0;
        for(int i = 0; i < 50; i++) {
            double res = 2*calculateIntegral(half_pi, 0, 1, n);
            std::cout << n << " " << res; 
            avg += res;
            if(i == 49){
                std::cout << " " << avg/50;
            }
            std::cout << std::endl;
        }
    }




    return 0;
}
