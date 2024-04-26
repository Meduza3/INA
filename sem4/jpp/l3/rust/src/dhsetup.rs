use rand::Rng;
use crate::gfint::{self, GFInt};

#[derive(Clone)]
pub struct DHSetup<const N: i32> {
    generator: gfint::GFInt<N>
}

impl<const N: i32> DHSetup<N> {

    fn power_mod(mut x: i64, mut y: i64, p: i64) -> i32 {
        let mut res = 1;
        x = x % p;
        while y > 0 {
            if y & 1 == 1 {
                res = (res * x) % p;
            }
            y = y >> 1;
            x = ( x * x) % p;
        }
        res as i32
    }

    fn prime_factors(mut n: i32) -> Vec<i32> {
        let mut factors: Vec<i32> = vec![];
        while n % 2 == 0 {
            factors.push(2);
            n /= 2;
        }
        let sqrt = (n as f64).sqrt() as i32;
        for i in (3..=sqrt).step_by(2) {
            while n % i32::from(i) == 0 {
                factors.push(i32::from(i));
                n /= i32::from(i);
            }
        }
        if n > 2 {
            factors.push(n);
        }
        factors
    }

    fn generate_generator() -> GFInt<N> {
        let generator: GFInt<N>;
        let mut rng = rand::thread_rng();
        let factors = Self::prime_factors(N - 1);
        loop {
            let candidate = rng.gen_range(1..=N - 1);
            let mut found = true;
            for factor in &factors {
                if Self::power_mod(candidate.into(), ((N - 1) / *factor).into(), N.into()) == 1 {
                    found = false;
                    break;
                }
            }
            if found {
                generator = gfint::GFInt::new(candidate);
                return generator;
            }
        }
    }

    pub fn new() -> Self {
        DHSetup {
            generator: Self::generate_generator()
        }
    }

    pub fn get_generator(self) -> GFInt<N> {
        self.generator
    }

    pub fn power(self, a: GFInt<N>, b: u128 ) -> GFInt<N> {
        let mut result: GFInt<N> = GFInt::new(1);
        for _ in 0..b {
            result = result * a;
        }
        result
    } 
}