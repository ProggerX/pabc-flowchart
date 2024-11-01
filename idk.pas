var
	a : Integer;
begin
	writeln('Give me a number');
	// Get a number
	readln(a);
	writeln('Your number is ', a, '!');
	if a mod 2 = 0 then writeln('wow! your num is even!');
	a += 2;
	writeln('Your num + 2 = ', a);
	for a := 0 to 10 step 2 do writeln(a);
	for a := 1 to 9 step 2 do writeln(a);
end.
