use std::ops::{Add, Sub, Mul, Div};
use std::fmt;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct GFInt<const N: i32> {
    value: i32,
}

impl<const N: i32> GFInt<N> {
    pub fn new(val: i32) -> Self {
        let value = val.rem_euclid(N);
        GFInt { value: value }
    }

    pub fn get_value(self) -> i32 {
        self.value
    }

    fn mod_inverse(a: i32, modulus: i32) -> i32 {
        let m0 = modulus;
        let mut x0 = 0;
        let mut x1 = 1;
        let mut a = a;
        let mut r#mod = modulus;

        if modulus == 1 {
            return 0;
        }

        while a > 1 {
            let q = a / r#mod;
            let t = r#mod;
            r#mod = a % r#mod;
            a = t;
            let t = x0;
            x0 = x1 - q * x0;
            x1 = t;
        }

        if x1 < 0 {
            x1 += m0;
        }

        x1
    }
}

impl<const N: i32> Add for GFInt<N> {
    type Output = Self;

    fn add(self, rhs: Self) -> Self::Output {
        Self::new((self.value + rhs.value).rem_euclid(N))
    }
}

impl<const N: i32> Sub for GFInt<N> {
    type Output = Self;

    fn sub(self, rhs:Self) -> Self::Output {
       Self::new((self.value - rhs.value).rem_euclid(N))
    }
}

impl<const N: i32> Mul for GFInt<N> {
    type Output = Self;

    fn mul(self, rhs: Self) -> Self::Output {
        Self::new((self.value as i64 * rhs.value as i64 % N as i64) as i32)
    }
}

impl<const N: i32> Div for GFInt<N> {
    type Output = Self;

    fn div(self, rhs: Self) -> Self::Output {
        if rhs.value == 0 {
            panic!("Attempt to divide by zero!");
        }
        let inverse = Self::mod_inverse(rhs.value, N);
        Self::new((self.value as i64 * inverse as i64 % N as i64) as i32)
    }
}

impl<const N: i32> fmt::Display for GFInt<N> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.value)
    }
}

fn main() {
    let a = GFInt::<1234567891>::new(4591);
    let b = GFInt::<1234567891>::new(1435);
    let c = GFInt::<1234567891>::new(5925);
    let d = GFInt::<1234567891>::new(14854);

    println!("a + b = {}", a + b);
    println!("a - d = {}", a - d);
    println!("d * c = {}", d * c);
    println!("c / b = {}", c / b);
    println!("is a equal to b*d? {}", if a == b * d { "Yes" } else { "No" });
}