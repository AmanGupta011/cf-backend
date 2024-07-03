#!/bin/bash
# Check if the correct number of arguments is provided
if [ "$#" -ne 9 ]; then
    echo "Usage: $0 <source1.cpp> <source2.cpp> <generator.cpp> <input.txt> <output1.txt> <output2.txt> <compiled1> <compiled2> <compiled_gen>"
    exit 1
fi

# Assign the arguments to meaningful variable names
source1=$1
source2=$2
generator=$3
input=$4
output1=$5
output2=$6
compiled1=$7
compiled2=$8
compiled_gen=$9

# Compile the source files
g++ "$source1" -o "$compiled1"
if [ $? -ne 0 ]; then
    echo "Error compiling $source1"
    exit 1
fi

g++ "$source2" -o "$compiled2"
if [ $? -ne 0 ]; then
    echo "Error compiling $source2"
    exit 1
fi

g++ "$generator" -o "$compiled_gen"
if [ $? -ne 0 ]; then
    echo "Error compiling $generator"
    exit 1
fi

# Run the tests in an infinite loop
for ((i = 1; ; ++i)); do
    "$compiled_gen" $i > "$input"
    "$compiled1" <"$input" > "$output1"
    "$compiled2" <"$input" > "$output2"
    if ! diff -w "$output1" "$output2"; then
        echo "Difference found in test $i"
        break
    fi
    echo "Passed test: $i"
done


# /Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-22-code.cpp /Users/amangupta/Desktop/cf-stress/backend/bin/contest/1746/b/solution.cpp /Users/amangupta/Desktop/cf-stress/backend/bin/contest/1746/b/generator.cpp /Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-22-input.txt /Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-22-output-participant.txt /Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-22-output-jury.txt /Users/amangupta/Desktop/cf-stress/backend/playground/contest/1746/b/sub_id-176323910-ticket-22-code /Users/amangupta/Desktop/cf-stress/backend/bin/contest/1746/b/solution /Users/amangupta/Desktop/cf-stress/backend/bin/contest/1746/b/generator