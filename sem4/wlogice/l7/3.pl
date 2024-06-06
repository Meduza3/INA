filozofowie :-
    mutex_create(W1),
    mutex_create(W2),
    mutex_create(W3),
    mutex_create(W4),
    mutex_create(W5),
    thread_create(filozof(1, W5, W1), _, [detached(true)]),
    thread_create(filozof(2, W1, W2), _, [detached(true)]),
    thread_create(filozof(3, W2, W3), _, [detached(true)]),
    thread_create(filozof(4, W3, W4), _, [detached(true)]),
    thread_create(filozof(5, W4, W5), _, [detached(true)]).

filozof(NUMER, LEWY, PRAWY) :-
    random(R),
    sleep(R),
    format('[~w] mysli~n',[NUMER]),
    format('[~w] chce prawy widelec~n',[NUMER]),
    mutex_lock(PRAWY),
    format('[~w] chce lewy widelec~n', [NUMER]),
    mutex_lock(LEWY),
    format('[~w] je~n', [NUMER]),
    format('[~w] odklada lewy widelec~n', [NUMER]),
    mutex_unlock(LEWY),
    format('[~w] odklada prawy widelec~n', [NUMER]),
    mutex_unlock(PRAWY),
    filozof(NUMER, LEWY, PRAWY).