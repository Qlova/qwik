#! /bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $DIR

function BasicTest {
	if [ "$?" -eq "1" ]; then
		echo -e "$1 \e[31mFAILED!\e[0m to compile D:"
		exit 1
	fi
	local OUTPUT=$(qwik run "$1.i" <<< $(echo -e "$3"))
	local DEFINED=$(echo -e "$2")
	if [ "$OUTPUT" = "$DEFINED" ]; then
		echo -e "$1 \e[32mPASSED!\e[0m"
	else
		echo -e "$1 \e[31mFAILED!\e[0m Got:"
		echo	 "$OUTPUT"
		echo "(Expecting)"
		echo	 "$DEFINED"
		exit 1
	fi
}

function GraphicsTest {
	echo -n "$1"
	echo -e " NOTSURE!"
}

function Passed {
	echo -e " \e[32mPASSED!\e[0m"
}

function FakeTest {
	echo -n "$1"
	echo -e " \e[32mFAILED!\e[0m"
}

function Failed {
	echo -e " \e[31mFAILED!\e[0m"
}

export -f BasicTest
export -f FakeTest
export -f GraphicsTest
export -f Passed
export -f Failed

if [[ ! -z "$1" ]]; then
	cd ./$1 && ./test.sh $2 $3 $4
else
	cd ../Basic && ./test.sh
fi
