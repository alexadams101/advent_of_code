import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

def part1(lines: Iterator[String]): String = 
    "Not Complete"
    
def part2(lines: Iterator[String]): String = 
    "Not Complete"