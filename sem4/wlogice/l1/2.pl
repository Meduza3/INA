on(Block1, Block2).

above(Block1, Block2) :- on(Block1, Block2).
above(Block1, Block2) :- on(Block1, X), above(X, Block2).