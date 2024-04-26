use crate::gfint::GFInt;
use crate::dhsetup::{self, DHSetup};
use rand::Rng;

#[derive(Clone)]
pub struct User<const N: i32> {
    setup: dhsetup::DHSetup<N>,
    secret: GFInt<N>,
    public_key: GFInt<N>,
    encryption_key: GFInt<N>,
}

impl<const N: i32> User<N> {

    fn generate_secret(&mut self) {
        let mut rng = rand::thread_rng();
        let secret_value: i32 = rng.gen_range(1..N-2);
        self.secret = GFInt::new(secret_value);
        self.public_key = self.setup.clone().power(self.setup.clone().get_generator(), self.secret.get_value() as u128)
    }

    pub fn get_public_key(&self) -> GFInt<N> {
        self.public_key
    }

    pub fn set_key(&mut self, a: GFInt<N>) {
        self.encryption_key = self.setup.clone().power(a, self.secret.get_value() as u128);
    }

    pub fn encrypt(&self, m: GFInt<N>) -> GFInt<N>{
       m * self.encryption_key
    }

    pub fn decrypt(&self, c: GFInt<N>) -> GFInt<N>{
        c / self.encryption_key
    }

    pub fn new(dh_setup: DHSetup<N>) -> Self {
        let mut user = User {
            setup: dh_setup,
            secret: GFInt::new(0),
            encryption_key: GFInt::new(0),
            public_key: GFInt::new(0)
        };
        user.generate_secret();
        user
    }
}