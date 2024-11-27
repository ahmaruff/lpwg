package main

import "fmt"

// Assign Integer to appropriate types
// the number should be store in the smallest type possible
// print integer on terminal

// 60
// -100
// 127
// 128
// -144243
// 255
// 256
// 144243
// 3641
// -4512
// 65535
// 1000000000000000000
// -65535
// -1000000000000000000
// -300000

func intTypes() {
    var a int8 = 60;
    var b int8 = -100;
    var c int8 = 127;
    var d int16 = 128;
    var e int32 = -144243;
    var f int16 = 255;
    var g int16 = 256;
    var h int32 = 144243;
    var i int16 = 3641;
    var j int16 = -4512;
    var k int32 = 65535;
    var l int64 = 1000000000000000000;
    var m int32 = -65535;
    var n int64 = -1000000000000000000;
    var o int32 = -300000;
    
    fmt.Println(a);
    fmt.Println(b);
    fmt.Println(c);
    fmt.Println(d);
    fmt.Println(e);
    fmt.Println(f);
    fmt.Println(g);
    fmt.Println(h);
    fmt.Println(i);
    fmt.Println(j);
    fmt.Println(k);
    fmt.Println(l);
    fmt.Println(m);
    fmt.Println(n);
    fmt.Println(o);
}
