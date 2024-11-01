var
	a : Integer;
begin
	writeln('Give me a number');
	// Get a number
	readln(a);
	writeln('Your number is ', a, '!');
	writeln('Now, give me a string');
	if a mod 2 = 0 then writeln('wow! your num is even!');
	a += 2;
	writeln('Your num + 2 = ', a);
end.
