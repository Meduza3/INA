use std::fmt;
use std::ops::{Add, AddAssign, Sub, SubAssign, Mul, MulAssign, Div, DivAssign};

struct GFInt<const CHARACTERISTIC: i32> {
    value: i32,
}

impl<const CHARACTERISTIC: i32> GFInt<CHARACTERISTIC> {
    fn new(val: i32) -> Self {
        GFInt {
            value: ((val % CHARACTERISTIC + CHARACTERISTIC) % CHARACTERISTIC),
        }
    }

    fn value(&self) -> i32 {
        self.value
    }

    fn mod_inverse(a: i32, m: i32) -> i32 {
        let (mut m0, mut t, mut q) = (m, 0, 0);
        let (mut x0, mut x1) = (0, 1);

        if m == 1 {
            return 0;
        }

        let mut a = a;
        while a > 1 {
            q = a / m;
            t = m;
            m = a % m;
            a = t;
            t = x0;
            x0 = x1 - q * x0;
            x1 = t;
        }

        if x1 < 0 {
            x1 += m0;
        }

        x1
    }
}

impl<const CHARACTERISTIC: i32> fmt::Display for GFInt<CHARACTERISTIC> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.value)
    }
}

impl<const CHARACTERISTIC: i32> Add for GFInt<CHARACTERISTIC> {
    type Output = Self;

    fn add(self, rhs: Self) -> Self::Output {
        Self::new(self.value + rhs.value)
    }
}

impl<const CHARACTERISTIC: i32> AddAssign for GFInt<CHARACTERISTIC> {
    fn add_assign(&mut self, rhs: Self) {
        self.value = (self.value + rhs.value) % CHARACTERISTIC;
    }
}

impl<const CHARACTERISTIC: i32> Sub for GFInt<CHARACTERISTIC> {
    type Output = Self;

    fn sub(self, rhs: Self) -> Self::Output {
        Self::new(self.value - rhs.value)
    }
}

impl<const CHARACTERISTIC: i32> SubAssign for GFInt<CHARACTERISTIC> {
    fn sub_assign(&mut self, rhs: Self) {
        self.value = (self.value - rhs.value + CHARACTERISTIC) % CHARACTERISTIC;
    }
}

impl<const CHARACTERISTIC: i32> Mul for GFInt<CHARACTERISTIC> {
    type Output = Self;

    fn mul(self, rhs: Self) -> Self::Output {
        let product = (self.value as i64 * rhs.value as i64) % CHARACTERISTIC as i64;
        Self::new(product as i32)
    }
}

impl<const CHARACTERISTIC: i32> MulAssign for GFInt<CHARACTERISTIC> {
    fn mul_assign(&mut self, rhs: Self) {
        self.value = ((self.value as i64 * rhs.value as i64) % CHARACTERISTIC as i64) as i32;
    }
}

impl<const CHARACTERISTIC: i32> Div for GFInt<CHARACTERISTIC> {
    type Output = Self;

    fn div(self, rhs: Self) -> Self::Output {
        if rhs.value == 0 {
            panic!("Attempt to divide by zero!");
        }
        let inverse = Self::mod_inverse(rhs.value, CHARACTERISTIC);
        let result = (self.value as i64 * inverse as i64) % CHARACTERISTIC as i64;
        Self::new(result as i32)
    }
}

impl<const CHARACTERISTIC: i32> DivAssign for GFInt<CHARACTERISTIC> {
    fn div_assign(&mut self, rhs: Self) {
        *self = *self / rhs;
    }
}