package com.example.demo

import java.util.ArrayList;
import java.util.HashMap;

fun main() {
    var names = ArrayList<String>()
    names.add("Alice")
    names.add("Bob")

    var ageMapping = HashMap<String, Int>()
    ageMapping["Alice"] = 30
    ageMapping["Bob"] = 25

    for (name in names) {
        println("$name is ${ageMapping.get(name)} years old.")
    }
}

