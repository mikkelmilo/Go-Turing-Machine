ATML (name subject to change) MACROS
(p_1, σ)M(p_2,p_3)
if M ends in h_a go to p_2
else if M ends in h_r go to p_3

------------------------------------------------
En TM starter i [_, {_},...]
dvs. (hs, a, _, b, >) → [_, b, {_},...]

Semantic rules:
TML programs must contain a start state, a reachable accept state, and a reachable reject state.
The same is true for macro definitions.

//TODO: update this description to also include hs
Macros are essentially self-contained, isolated Turing Machines within the "main" Turing Machine.
Thus, just like the main TM, inside a macro definition there must exist a start-, accept-, and reject state,
labeled hs, ha, and hr, respectively.
The syntax for "calling" a macro is
	(p_1, σ)M(p_2,p_3)
The meaning is as follows:
If the TM is in state p_1 and the tape head reads the symbol σ, execute the macro by going to the state 'hs'
inside the macro definition. If the execution of the macro accepts, go to state p_2, else if it rejects go to state p_3.
The TM defined by the macro is isolated. This means it cannot transition to states defined outside the macro.
The compiler handles this by overshadowing all state names inside a macro. Thus, if you try to transition to a state
defined outside a macro, the compiler will not complain because it simply generates a new state that is local to this macro.


How Macro Substitution Works:

Example:
given a macro definition:
define swap {
	(hs, ha, 0, 1, _)
	(hs, hr, 1, 0, _)
}
and following usage:
(hs, a, _, 0, _)
(a, _)swap(c,d)
(c,ha,_,_,_)
(d,hr,_,_,_)

This program will write a 0 and go to state a. Then it will call the macro 'swap'
which is supposed to swap the value of the tape element. It will then go to the state c if a 0 was swapped to 1,
or it will go to state d if a 1 was swapped to a zero. In state c it will go to the accept state,
and in state d it will go to the reject state.

//TODO: should the entire macro be copied for each application? or can it be prevented?
The inner working of the substitution is as follows:
for each macro application (p_1, σ)M(p_2,p_3) a new unique instance of the macro is added to the TML program,
where the states 'hs', 'ha', and 'hr' of the macro are renamed to 'hsM_i', 'haM_i', 'hrM_i', respectively,
to avoid conflict. The i is a unique number generated for each macro application.
The following transition rules are added to the TML program:
	(p_1, hsM_i, σ, σ, _)
	(haM_i, p_1, _, _, _)
	(hsM_i, p_2, _, _, _)
The first transition rule ensures that this macro instance is reachable, and its start state is reached whenever
the TM is in state p_1 and sees symbol σ.
The last two transition rules make the transitions from the macro's accept state to p_1,
and from the macros reject state to p_2. Both of these transitions are unconditional, so we don't care about what the tape
head contains, and we don't move the tape head either.

## TML quirks
start, accept, and reject states cannot *only* exist in macro application commands. They must also exist in "normal" commands.
Otherwise the compiler will (albeit somewhat errornously) report an error that this state is missing.