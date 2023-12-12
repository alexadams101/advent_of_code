import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

class Game(var number: Int, var red: Int, var green: Int, var blue: Int)

def part1(lines: Iterator[String]): String = 
    def function(g: Game): Boolean = 
        val redMax = 12
        val greenMax = 13
        val blueMax = 14
        redMax >= g.red && greenMax >= g.green && blueMax >= g.blue

    lines.map(toGame(_)).filter(function(_)).map(_.number).sum.toString()
    
def part2(lines: Iterator[String]): String = 
    def function(g: Game): Int = 
        g.red * g.green * g.blue

    lines.map(toGame(_)).map(function(_)).sum.toString()

def toGame(s: String): Game =
    def toColour(s: String): (String, Int) =
        val strings = s.substring(1).split(" ")
        (strings(1), strings(0).toInt)

    val strings = s.split(":")
    val game = strings(0).split(" ")(1).toInt
    val colours = strings(1).split(",|;").map(toColour(_))
    val red = colours.filter(_._1.equals("red")).map(_._2).max
    val green = colours.filter(_._1.equals("green")).map(_._2).max
    val blue = colours.filter(_._1.equals("blue")).map(_._2).max
    new Game(game, red, green, blue)