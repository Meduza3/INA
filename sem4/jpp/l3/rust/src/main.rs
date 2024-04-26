use dhsetup::DHSetup;

use crate::gfint::GFInt;

use crate::user::User;

mod gfint;
mod dhsetup;
mod user;


fn main() {
    const CHARACTERISTIC: i128 = 1234567891;
    let dh_setup: DHSetup<CHARACTERISTIC> = DHSetup::new();

    let mut alice: User<CHARACTERISTIC> = User::new(dh_setup.clone());
    let mut bob: User<CHARACTERISTIC> = User::new(dh_setup.clone());

    alice.set_key(bob.get_public_key());
    bob.set_key(alice.get_public_key());

    let message: GFInt<CHARACTERISTIC> = GFInt::new(6969420);
    let encrypted_message = alice.encrypt(message);
    let decrypted_message = bob.decrypt(encrypted_message);

    println!("Wiadomość oryginalna: {}", message);
    println!("Wiadomość zaszyfrowana: {}", encrypted_message);
    println!("Wiadomość odszyfrowana: {}", decrypted_message);
}
