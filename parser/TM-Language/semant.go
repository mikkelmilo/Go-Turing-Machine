package parser

/*
 * This file contains a Tree walker which checks the semantics of the program.
 * This includes:
 *    - the program contains hs, ha, and hr states
 *    - all states must be reachable from hs (the same is true for macros and their states)
 *    - macros must also contain hs, ha, and hr states
 *    - warnings will be produced (if enabled) if there are unbreakable cycles, ie. 
 *      cycles which has no sequence of transitions to either ha or hr 
 */
