data class Student(val name: String, val percentage: Double, val section: String)

fun main() {

    println("List Operations")

    val names = listOf("Yuvraj", "Manshi", "Anju", "Muskan")
    println("Read-only List: $names")
    println("Contains Anju: ${names.contains("Anju")}")
    println("Element at index 1: ${names[1]}")

    val mutableNames = mutableListOf("Yuvraj", "Manshi", "Anju")
    mutableNames.add("Muskan")
    mutableNames.remove("Manshi")
    println("Mutable List: $mutableNames")

    println("\nSet Operations")

    val fruits = setOf("Apple", "Mango", "Banana")
    println("Read-only Set: $fruits")
    println("Contains Mango: ${fruits.contains("Mango")}")

    val mutableFruits = mutableSetOf("Apple", "Mango", "Banana")
    mutableFruits.add("Orange")
    mutableFruits.remove("Apple")
    println("Mutable Set: $mutableFruits")

    println("\nMap Operations")

    val marks = mapOf("Yuvraj" to 85, "Manshi" to 90, "Anju" to 88)
    println("Read-only Map: $marks")
    println("Yuvraj's Marks: ${marks["Yuvraj"]}")
    println("Contains Anju: ${marks.containsKey("Anju")}")

    val mutableMarks = mutableMapOf("Yuvraj" to 85, "Manshi" to 90)
    mutableMarks["Anju"] = 88
    mutableMarks.remove("Manshi")
    println("Mutable Map: $mutableMarks")

    println("\nMap, Filter and Reduce Operations")

    val numbers = listOf(10, 20, 30, 40, 50)
    println("Original List: $numbers")

    val doubled = numbers.map { it * 2 }
    println("After map: $doubled")

    val filtered = numbers.filter { it > 20 }
    println("After filter: $filtered")

    val total = numbers.reduce { a, b -> a + b }
    println("After reduce: $total")

    val upperNames = names.map { it.uppercase() }
    println("Uppercase Names: $upperNames")

    val longNames = names.filter { it.length > 4 }
    println("Names with more than 4 letters: $longNames")

    println("\nSorting Students")

    val students = listOf(
        Student("Yuvraj", 85.5, "A"),
        Student("Manshi", 92.0, "B"),
        Student("Anju", 78.5, "A"),
        Student("Muskan", 89.0, "B"),
        Student("Lalit", 75.0, "A")
    )

    println("Original Student List:")
    students.forEach { println(it) }

    val sortedStudents = students.sortedByDescending { it.percentage }

    println("\nStudents Sorted by Percentage:")
    sortedStudents.forEach { println(it) }

    println("\nGrouping Students")

    val groupedStudents = students.groupBy { it.section }

    groupedStudents.forEach { (section, studentList) ->
        println("Section $section:")
        studentList.forEach { println(it.name) }
    }
}