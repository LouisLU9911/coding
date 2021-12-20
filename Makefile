format:
	find ./ -iname "*.h" -o -iname "*.cc" | xargs clang-format -sort-includes -i
	buildifier -r .

build:
	bazel build //...

clean:
	bazel clean

