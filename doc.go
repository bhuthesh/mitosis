// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Mitosis allows Go applications to easily fork themselves while preserving
arbitrary application state and inherit file descriptors.
In this context, 'arbitrary' means: anything you can marshal into a byte slice.
*/
package mitosis
