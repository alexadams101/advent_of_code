import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

def part1(lines: Iterator[String]): String = 
    lines.map(toNumbers).map(getNextNumberPart1).sum.toString()
    
def part2(lines: Iterator[String]): String = 
    lines.map(toNumbers).map(getNextNumberPart2).sum.toString()

def toNumbers(s: String): List[Int] =
    s.split(" ").map(_.toInt).toList

def getNextNumberPart1(ints: List[Int]): Int =
    if (ints.filter(_ != 0).length == 0)
        return 0
    var nextArray = List[Int]()
    for (i <- 0 until ints.length - 1)
        nextArray = nextArray:+ ints(i + 1) - ints(i) 
    val nextNumber = getNextNumberPart1(nextArray)
    ints(ints.length - 1) + nextNumber

def getNextNumberPart2(ints: List[Int]): Int =
    if (ints.filter(_ != 0).length == 0)
        return 0
    var nextArray = List[Int]()
    for (i <- 0 until ints.length - 1)
        nextArray = nextArray:+ ints(i + 1) - ints(i) 
    val nextNumber = getNextNumberPart2(nextArray)
    ints(0) - nextNumber
    