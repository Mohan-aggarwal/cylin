@seed = global i32 0

declare i32 @abs(i32 %x)

declare void @printi(i32 %x)

declare i32 @utime()

define i32 @rand() {
0:
	%1 = load i32, i32* @seed
	%2 = mul i32 %1, 22695477
	%3 = add i32 %2, 1
	store i32 %3, i32* @seed
	%4 = call i32 @abs(i32 %3)
	call void @printi(i32 %4)
	ret i32 %4
}

define i32 @main() {
entry:
	%0 = call i32 @utime()
	store i32 %0, i32* @seed
	%1 = call i32 @rand()
	ret i32 0
}

