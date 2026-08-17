 #!/usr/bin/env sh

set -o nounset # – падай на необъявленных переменных
set -o errexit # – падай на первой ошибке
set -o pipefail # – падай на проблемах в pipes

MY_VAR='Hi!'

echo $MY_VAR

if [ -z "$MY_VAR" ]; then
    echo "MY_VAR is empty"
else
    echo "MY_VAR is not empty"
fi


MY_VAR=1

test "$MY_VAR" -eq 1 \
&& echo 'MY_VAR is 1!' \
|| echo 'MY_VAR is not 1 :('


SUM_RESULT=0

while [ "$SUM_RESULT" -le 5 ]; do
    echo "Welcome $SUM_RESULT times"
    SUM_RESULT=$(( SUM_RESULT + 1 ))
done

MY_ARRAY=(1 2 3 4 5 6 7 7 7)

echo "Array size: ${#MY_ARRAY[*]}"

for item in ${MY_ARRAY[*]}; do
 echo "item: $item"
done

while read -r line; do 
    echo "line: $line"
done < items.txt


my_ls() { # declaring
  echo 'Listing files'
  ls
  echo 'Done!'
}

my_ls


my_function() {
 local variable=1
 readonly OTHER_VARIABLE=2
}

my_function

echo "OTHER_VARIABLE is $OTHER_VARIABLE"
echo "variable is: $variable"


sum_two_numbers() {
 local first
 local second

 first="$1"
 second="$2"

 echo "$((first + second))"
}

sum_two_numbers 2 3 


greet() {
 local user
 user="${1:-pal}"

 echo "Hello, $user"
}

greet sobolevn
greet

echo_all_words() {
  echo "Words: $*"
}

echo_all_words 'cat' 'dog' 'piggy'

echo_all_words2() {
  Words="$@"

  for item in ${Words[*]}; do
    echo "Word: $item"
  done
}

echo_all_words2 'cat' 'dog' 'piggy'


can_fail() {
 if [ "$1" -eq 5 ]; then
 echo 'error'
 return 1
 fi
 
 echo 'correct'
 return 0
}
 
can_fail 1
echo "status code is: $?"

can_fail 5
echo "status code is: $?"

can_fail2() {
 if [ "$1" -eq 5 ]; then
 echo 'error'
 else
 echo 'correct'
 fi
 }
## subshell
 first=$(can_fail2 1)
 echo "returns: $first code: $?"

 second=$(can_fail2 5)
 echo "returns: $second code: $?"