
data class Student(
    val id: Int,
    val name: String,
    val email: String?
)

data class Course(
    val course_Id: Int,
    val course_Name: String,
    val fee: Double,
    val instructor: String?
)

data class Registration(
    val student: Student,
    val course: Course,
    val marks: Int?
)

fun Registration.hasPassed(): Boolean {
    return (marks ?: 0) >= 40
}

fun Course.FeeafterDiscount(discount: Double): Double {
    return fee - (fee * discount / 100)
}

class CourseReport {
    private val reportLines = mutableListOf<String>()

    fun heading(title: String) {
        reportLines.add("\n$title")
    }

    fun details(message: String) {
        reportLines.add(message)
    }

    fun show() {
        reportLines.forEach { println(it) }
    }
}

fun courseReport(block: CourseReport.() -> Unit) {
    val report = CourseReport()
    report.block()
    report.show()
}


fun main() {

    val course1 = Course(101, "Kotlin Programming", 32999.0, "Mr. Hemraj")
    val course2 = Course(102, "Java Programming", 45670.0, "Mr. Manish")
    val course3 = Course(103, "Python Programming", 47429.0, null)

    val student1 = Student(1, "Yuvraj", "yuvraj@gmail.com")
    val student2 = Student(2, "Manshi", "manshi@gmail.com")
    val student3 = Student(3, "Anju", null)

    val registrations = listOf(
        Registration(student1, course1, 85),
        Registration(student2, course2, 75),
        Registration(student3, course1, 35),
        Registration(student1, course3, 90),
        Registration(student2, course3, null)
    )

    println("COURSE REGISTRATION SYSTEM")
    println("Instructor: ${course3.instructor ?: "Not Assigned"}")
    println("Student Email: ${student3.email ?: "Not Provided"}")

    val studentNames = registrations.map { it.student.name }
    println("\nStudent Names: $studentNames")

    val passedStudents = registrations.filter { it.hasPassed() }
    println("\nStudents Who Passed:")
    passedStudents.forEach {
        println("${it.student.name} - ${it.course.course_Name}")
    }

    val totalFees = registrations
        .map { it.course.fee }
        .reduce { total, fee -> total + fee }

    println("\nTotal Registration Fees: Rs. $totalFees")

    val groupedCourses = registrations.groupBy {
        it.course.course_Name
    }

    println("\nRegistrations Grouped By Course:")
    groupedCourses.forEach { (course, students) ->
        println("$course: ${students.size} registrations")
    }

    println("\nDiscounted Course Fee:")
    println("${course1.course_Name}: Rs. ${course1.FeeafterDiscount(10.0)}")

    courseReport {
        heading("COURSE REGISTRATION REPORT")

        details("Total Registrations: ${registrations.size}")

        details("Total Fees Collected: Rs. $totalFees")

        details("Passed Registrations: ${passedStudents.size}")

        details("Failed Registrations: ${
            registrations.count { !it.hasPassed() }
        }")

        heading("COURSE-WISE DETAILS")

        groupedCourses.forEach { (course, students) ->
            details("$course -> ${students.size} students")
        }

        heading("REPORT COMPLETED!!!!!!!!!!!!!!")
    }
}
