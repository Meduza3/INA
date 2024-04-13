use std::ops::{Add, Sub, Mul, Div};
use std::fmt;

const CHARACTERISTIC: i32 = 1234577;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct GF1234577Int {
    value: i32,
}

impl GF1234577Int {
    fn new(val: i32) -> Self {
        GF1234577Int {value: val.rem_euclid(CHARACTERISTIC) }
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

impl Add for GF1234577Int {
    type Output = GF1234577Int;

    fn add(self, rhs: Self) -> Self::Output {
        GF1234577Int::new((self.value + rhs.value).rem_euclid(CHARACTERISTIC))
    }
}

impl Sub for GF1234577Int {
    type Output = GF1234577Int;

    fn sub(self, rhs:Self) -> Self::Output {
        GF1234577Int::new((self.value - rhs.value).rem_euclid(CHARACTERISTIC))
    }
}

impl Mul for GF1234577Int {
    type Output = GF1234577Int;

    fn mul(self, rhs: Self) -> Self::Output {
        GF1234577Int::new((self.value as i64 * rhs.value as i64 % CHARACTERISTIC as i64) as i32)
    }
}

impl Div for GF1234577Int {
    type Output = GF1234577Int;

    fn div(self, rhs: Self) -> Self::Output {
        if rhs.value == 0 {
            panic!("Attempt to divide by zero!");
        }
        let inverse = Self::mod_inverse(rhs.value, CHARACTERISTIC);
        GF1234577Int::new((self.value as i64 * inverse as i64 % CHARACTERISTIC as i64) as i32)
    }
}

impl fmt::Display for GF1234577Int {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.value)
    }
}

fn main() {
    let a = GF1234577Int::new(4591);
    let b = GF1234577Int::new(1435);
    let c = GF1234577Int::new(5925);
    let d = GF1234577Int::new(14854);

    println!("a + b = {}", a + b);
    println!("a - d = {}", a - d);
    println!("d * c = {}", d * c);
    println!("c / b = {}", c / b);
    println!("is a equal to b*d? {}", if a == b * d { "Yes" } else { "No" });
}
