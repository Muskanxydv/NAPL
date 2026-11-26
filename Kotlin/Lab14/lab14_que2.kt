fun main() {

    println("Nullable and Non-Nullable Variables")

    var name: String = "Yuvraj"
    var surname: String? = "Kumar"
    var city: String? = null

    println("Non-nullable Name: $name")
    println("Nullable Surname: $surname")
    println("Nullable City: $city")

    println("\nSafe Call Operator")

    val studentName: String? = "Manshi"
    println("Name Length: ${studentName?.length}")

    val anotherName: String? = null
    println("Name Length: ${anotherName?.length}")

    println("\nElvis Operator")

    val student: String? = null
    val result = student ?: "Anju"

    println("Student Name: $result")

    val student2: String? = "Muskan"
    val result2 = student2 ?: "Lalit"

    println("Student Name: $result2")

    println("\nNot-null Assertion Operator")

    val name1: String? = null

    try {
        println(name1!!)
    } catch (e: NullPointerException) {
        println("NullPointerException caught successfully")
    }

    println("Program continues normally")
}