import scala.io.Source

object App {
    def main(args: Array[String]) = {
        val filename = "day_1/data.txt"
        for (line <- Source.fromFile(filename).getLines) {
            println(line)
        }
    }
}