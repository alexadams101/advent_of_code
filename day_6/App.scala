import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

def part1(lines: Iterator[String]): String = 
    val races = mapToRaces(lines)
    races.map(calculateWaysToWin).product.toString()
    
def part2(lines: Iterator[String]): String = 
    val race = mapToRace(lines)
    calculateWaysToWin(race).toString()

def calculateWaysToWin(race: (Long, Long)): Int =
    val duration = race._1
    val record = race._2
    var q = 0
    for (i <- 0.toLong until duration) {
        if (i * (duration-i) > record) {
            q = q + 1
        }
    }
    q

def calculateWaysToWin(race: (Long, Long), range: Range): Int =
    val duration = race._1
    val record = race._2
    var q = 0
    for (i <- range) {
        if (i * (duration-i) > record) {
            q = q + 1
        }
    }
    q

def mapToRaces(lines: Iterator[String]): List[(Long, Long)] = 
    val numbers = lines.map(_.split(" ").filter(_.nonEmpty).drop(1)).toList
    var list = List[(Long, Long)]()
    for (i <- 0 until numbers(0).length) {
        list = list:+ (numbers(0)(i).toLong, numbers(1)(i).toLong)
    }
    list

def mapToRace(lines: Iterator[String]): (Long, Long) = 
    val numbers = lines.map(_.split(":")(1).replace(" ", "")).toList
    (numbers(0).toLong, numbers(1).toLong)