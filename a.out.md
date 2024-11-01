flowchart TB
var0[Объявить a типа Integer]
mbbeg([Начало блока])
mb0["writeln('Give me a number')"]
mbbeg-->mb0
mb1["readln(a)"]
mb0-->mb1
mb2["writeln('Your number is ', a, '!')"]
mb1-->mb2
mb3["writeln('Now, give me a string')"]
mb2-->mb3
mb4if{"a mod 2 = 0 ?"}
mb4ifthen["writeln('wow! your num is even!')"]
mb4if-->|тогда|mb4ifthen
mb4if-->|иначе|mb4ifend
mb4ifthen-->mb4ifend
mb4ifend[Конец условия]
mb3-->mb4if
mb5as["Присвоить a значение a + 2"]
mb4ifend-->mb5as
mb6["writeln('Your num + 2 = ', a)"]
mb5as-->mb6
mbend([Конец блока])
mb6-->mbend