template <typename T>
class USer {

    private:
        DHSetup setup;
        T secret;
        T privateKey;

    public:
        User(DHSetup& setup){
            this->setup = setup;
        }

        T getPublicKey(){
            return setup.power(setup.getGenerator(), secret);
        }

        void setKey(T a){
            setup.power(a, secret);
        }

        T encrypt(T m){
            return setup.power(m, privateKey);
        }

        T decrypt(T c){
            return x/privateKey;
        }
};