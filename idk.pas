var
	a : Integer;
	b : String;
begin
	writeln('Give me a number');
	// Get a number
	readln(a);
	writeln('Your number is ', a, '!');
	writeln('Now, give me a string');
	// Get a string
	readln(b);
	writeln('Your string is ', b, '!');
	if a mod 2 = 0 then begin
		writeln('wow! your num is even!');
	end;
end.
