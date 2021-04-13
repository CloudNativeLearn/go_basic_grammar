package main

function test (){
var num = []
var i

for (i = 0; i < 10; i++) {
num[i] = function (i) {
console.log(i)
}(i)
}
return num[0]
}

test
