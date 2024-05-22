#include <iostream>
#include <mutex>
#include <chrono>
#include <thread>
#include <random>

const int ilosc_filozofow = 5;
const int min_dobranoc = 50;
const int max_dobranoc = 100;
const int ilosc_posilkow = 5;

int main() {
  std::vector<std::thread> filozofowie(ilosc_filozofow);
  std::vector<Widelec> widelce(ilosc_filozofow);

  for(int i = 0; i < ilosc_filozofow; i++) {
    Widelec left = widelce[i];
    Widelec right = widelce[(i + 1) % ilosc_filozofow];

  }
}

class Filozof {
  private:
    int id;
    Widelec* left;
    Widelec* right;
    int ilosc_posilkow;
    
  public:
    Filozof(int id, Widelec* left, Widelec* right, int ilosc_posilkow);


    void operator()() {
      this->start();
    }

    void mysl() {
      std::cout << "Filozof " << this->id << " mysli \n";
      this->dobranoc();
      std::cout << "Filozof " << this->id << " skonczyl myslec \n";
    }

    void jedz() {
      std::cout << "Filozof " << this->id << " je \n";
      this->dobranoc();
      std::cout << "Filozof " << this->id << " skonczyl jesc \n";
    } 

    void podnies_widelec(Widelec* widelec) {
      widelec->mutex.lock();
    }

    void odloz_widelec(Widelec* widelec) {
      widelec->mutex.unlock();
    }

    void dobranoc() {
      std::random_device rd;
      std::mt19937 rng(rd());
      std::uniform_int_distribution<int> uni(min_dobranoc, max_dobranoc);
      std::this_thread::sleep_for(std::chrono::milliseconds((uni(rng) + max_dobranoc)));
    }

    void start() {
      for(int i = 0; i < this->ilosc_posilkow; i++){
        this->mysl();
        this->podnies_widelec(this->left);
        this->podnies_widelec(this->right);
        this->jedz();
        this->odloz_widelec(this->left);
        this->odloz_widelec(this->right);
      }
    }
};

class Widelec {
  public:
    std::mutex mutex;
};