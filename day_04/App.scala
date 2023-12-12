import scala.io.Source

@main def part1(filename: String): Unit = println(s"Answer to part 1 is: ${part1(loadData(filename))}")

@main def part2(filename: String): Unit = println(s"Answer to part 2 is: ${part2(loadData(filename))}")

def loadData(filename: String): Iterator[String] = Source.fromFile(filename).getLines

class Game(var opp: List[Int], var self: List[Int])

def part1(lines: Iterator[String]): String = 
    lines.map(toGame(_)).map(getMatchingValues(_)).map(toPower(_)).sum.toString()
    
def part2(lines: Iterator[String]): String = 
    val values = lines.map(toGame(_)).map(getMatchingValues(_)).toList
    getTotalScratchCards(values).toString()

def toGame(s: String): Game =
    val vals = s.split(":")(1).split("\\|")
    val opp = vals(0).drop(1).dropRight(1).split(" ").filter(_.nonEmpty).map(_.toInt).toList
    val self = vals(1).drop(1).split(" ").filter(_.nonEmpty).map(_.toInt).toList
    new Game(opp, self)

def getMatchingValues(g: Game): List[Int] =
    g.opp.intersect(g.self)

def toPower(i: List[Int]): Int = 
    if (i.size > 0) 
        Math.pow(2, i.size-1).toInt
    else
        0
    
def getTotalScratchCards(i: List[List[Int]]): Int = 
    var counts = Array.fill(i.size)(1)
    val sizes = i.map(_.size)
    var total = 0

    for (n <- 0 until i.size) {
        if (sizes(n)>0) {
            val no = sizes(n) * counts(n)
            total = total + no
            val e = no / sizes(n)
            for (l <- n+1 to n+sizes(n)) {
                counts.update(l, counts(l) + e)
            }
        }
    }
    
    counts.sum
