with Interfaces.C;

package Main is
	type Diophantine_Solution is record
		X : Interfaces.C.int;
		Y : Interfaces.C.int;
	end record;
	
	function silnia_r(N : Interfaces.C.int) return Interfaces.C.int;
	pragma Import (C, silnia_r, "silnia_r");
	
	function nwd_r(A, B : Interfaces.C.int) return Interfaces.C.int;
	pragma Import (C, nwd_r, "nwd_r");
	
	function extended_euclid_r(A, B : Interfaces.C.int) return Diophantine_Solution;
	pragma Import (C, extended_euclid_r, "extended_euclid_r");

	function silnia_l(N : Interfaces.C.int) return Interfaces.C.int;
	pragma Import (C, silnia_l, "silnia_l");
	
	function nwd_l(A, B : Interfaces.C.int) return Interfaces.C.int;
	pragma Import (C, nwd_l, "nwd_l");
	
	function extended_euclid_l(A, B : Interfaces.C.int) return Diophantine_Solution;
	pragma Import (C, extended_euclid_l, "extended_euclid_l");
end Main;