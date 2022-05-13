build: format
	bazel build //:all_binaries

format:
	find ./ -iname "*.h" -o -iname "*.cc" | xargs clang-format -sort-includes -i
	buildifier -r .
	yapf -i -r .

clean:
	bazel clean

